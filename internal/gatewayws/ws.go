package gatewayws

import (
	"bytes"
	"context"
	"io"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"cdr.dev/slog"
	"github.com/coadler/played"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/etcdserverpb"
	"github.com/etcd-io/etcd/clientv3/concurrency"
	"github.com/go-redis/redis"
	"golang.org/x/xerrors"
	"nhooyr.io/websocket"

	"github.com/tatsuworks/czlib"
	"github.com/tatsuworks/gateway/discordetf"
	"github.com/tatsuworks/gateway/handler"
)

const (
	IdentifyMutexRootName = "/gateway/identify/"
	GatewayETF            = "wss://gateway.discord.gg/?v=6&encoding=etf&compress=zlib-stream"
)

type Session struct {
	ctx    context.Context
	cancel func()
	wg     *sync.WaitGroup

	log slog.Logger

	token   string
	shardID int
	shards  int

	seq    int64
	sessID string
	last   int64

	wsConn *websocket.Conn
	zr     io.ReadCloser

	interval time.Duration
	trace    string

	lastHB  time.Time
	lastAck time.Time

	buf  *bytes.Buffer
	hbuf *bytes.Buffer

	etcd       *clientv3.Client
	etcdSess   *concurrency.Session
	identifyMu *concurrency.Mutex

	state  *handler.Client
	rc     *redis.Client
	played *played.Client
}

func NewSession(
	logger slog.Logger,
	wg *sync.WaitGroup,
	rdb *redis.Client,
	etcdCli *clientv3.Client,
	token string,
	shardID, shards int,
) (*Session, error) {
	c, err := handler.NewClient()
	if err != nil {
		return nil, xerrors.Errorf("create state handler: %w", err)
	}

	sess := &Session{
		wg:      wg,
		log:     logger.With(slog.F("shard", shardID)),
		token:   token,
		shardID: shardID,
		shards:  shards,

		// start with a 1kb buffer
		buf:  bytes.NewBuffer(make([]byte, 0, 1<<10)),
		hbuf: bytes.NewBuffer(nil),

		etcd: etcdCli,

		state: c,
		rc:    rdb,
	}

	sess.loadSessID()
	sess.loadSeq()
	return sess, nil
}

func (s *Session) initEtcd() error {
	sess, err := concurrency.NewSession(s.etcd, concurrency.WithContext(s.ctx), concurrency.WithTTL(20))
	if err != nil {
		return xerrors.Errorf("get etcd session: %w", err)
	}

	s.etcdSess = sess
	s.identifyMu = concurrency.NewMutex(sess, IdentifyMutexRootName+strconv.Itoa(s.shardID%16))
	return nil
}

func (s *Session) shouldResume() bool {
	return s.seq != 0 && s.sessID != ""
}

func (s *Session) Open(ctx context.Context, token string, playedAddr string) error {
	s.wg.Add(1)
	defer s.wg.Done()

	s.ctx, s.cancel = context.WithCancel(ctx)
	defer s.cancel()

	played, err := played.NewClient(s.ctx, playedAddr)
	if err != nil {
		return xerrors.Errorf("connect to played: %w", err)
	}
	s.played = played

	s.lastAck = time.Time{}

	err = s.initEtcd()
	if err != nil {
		return err
	}

	// only acquire the identify lock if we know we won't send a resume
	if !s.shouldResume() {
		s.log.Debug(s.ctx, "acquiring lock, no ability to resume")
		err = s.acquireIdentifyLock()
		if err != nil {
			return xerrors.Errorf("grab identify lock: %w", err)
		}
		s.log.Debug(s.ctx, "lock acquired")

	} else {
		s.log.Debug(s.ctx, "skipping lock, attempting resume", slog.F("sess", s.sessID), slog.F("seq", s.seq))
	}

	r, err := czlib.NewReader(bytes.NewReader(nil))
	if err != nil {
		return xerrors.Errorf("initialize zlib: %w", err)
	}
	s.zr = r
	defer r.Close()

	c, _, err := websocket.Dial(s.ctx, GatewayETF, nil)
	if err != nil {
		return xerrors.Errorf("dial gateway: %w", err)
	}
	s.wsConn = c
	s.wsConn.SetReadLimit(512 << 20)

	err = s.readHello()
	if err != nil {
		return xerrors.Errorf("handle hello message: %w", err)
	}

	if s.shouldResume() {
		s.log.Info(s.ctx, "sending resume")
		err := s.writeResume()
		if err != nil {
			return xerrors.Errorf("send resume: %w", err)
		}
	} else {
		s.last = 0
		s.log.Info(s.ctx, "sending identify")
		err := s.writeIdentify()
		if err != nil {
			return xerrors.Errorf("send identify: %w", err)
		}
	}

	go s.sendHeartbeats()
	go s.logTotalEvents()
	// go s.rotateStatuses()

	s.log.Info(s.ctx, "websocket connected, waiting for events")

	for {
		err = s.readMessage()
		if err != nil {
			var werr websocket.CloseError
			if xerrors.As(err, &werr) {
				if werr.Code == 4006 {
					s.seq = 0
					s.persistSeq()
					s.sessID = ""
					s.persistSessID()
				}
			}

			err = xerrors.Errorf("read message: %w", err)
			break
		}

		var ev *discordetf.Event
		ev, err = discordetf.DecodeT(s.buf.Bytes())
		if err != nil {
			err = xerrors.Errorf("decode event: %w", err)
			break
		}

		if ev.S != 0 {
			atomic.StoreInt64(&s.seq, ev.S)
		}

		if handled, err := s.handleInternalEvent(ev); handled {
			if err != nil {
				return err
			}

			continue
		}

		if ev.T == "PRESENCE_UPDATE" {
			err := s.played.WritePresence(s.ctx, ev.D)
			if err != nil {
				s.log.Error(s.ctx, "send played event", slog.Error(err))
			}
			continue
		}

		// this is jank, will fix soon
		var requestMembers int64
		requestMembers, err = s.state.HandleEvent(ev)
		if err != nil {
			s.log.Error(s.ctx, "handle state event", slog.Error(err))
			continue
		}

		err = s.rc.RPush("gateway:events:"+ev.T, ev.D).Err()
		if err != nil {
			s.log.Error(s.ctx, "push event to redis", slog.Error(err))
		}

		if requestMembers != 0 {
			s.log.Debug(s.ctx, "requesting guild members", slog.F("guild", requestMembers))
			err := s.requestGuildMembers(requestMembers)
			if err != nil {
				s.log.Error(s.ctx, "request guild members", slog.Error(err))
			}
		}
	}

	s.persistSeq()
	_ = c.Close(websocket.StatusNormalClosure, "")
	s.log.Info(s.ctx, "closed")
	return err
}

func (s *Session) persistSeq() {
	err := s.rc.Set(s.fmtSeqKey(), s.seq, 0).Err()
	if err != nil && !xerrors.Is(err, redis.Nil) {
		s.log.Error(s.ctx, "save seq", slog.Error(err))
	}
}

func (s *Session) loadSeq() {
	sess, err := s.rc.Get(s.fmtSeqKey()).Result()
	if err != nil && !xerrors.Is(err, redis.Nil) {
		s.log.Error(s.ctx, "load session id", slog.Error(err))
	}

	if sess == "" {
		return
	}

	s.seq, err = strconv.ParseInt(sess, 10, 64)
	if err != nil {
		s.log.Error(s.ctx, "parse session id", slog.Error(err))
	}
}

func (s *Session) persistSessID() {
	err := s.rc.Set(s.fmtSessIDKey(), s.sessID, 0).Err()
	if err != nil && !xerrors.Is(err, redis.Nil) {
		s.log.Error(s.ctx, "save seq", slog.Error(err))
	}
}

func (s *Session) loadSessID() {
	sess, err := s.rc.Get(s.fmtSessIDKey()).Result()
	if err != nil && !xerrors.Is(err, redis.Nil) {
		s.log.Error(s.ctx, "load session id", slog.Error(err))
	}

	s.sessID = sess
}

func (s *Session) fmtSeqKey() string {
	return "gateway:seq:" + strconv.Itoa(s.shardID)
}

func (s *Session) fmtSessIDKey() string {
	return "gateway:sess:" + strconv.Itoa(s.shardID)
}

func (s *Session) handleInternalEvent(ev *discordetf.Event) (bool, error) {
	switch ev.Op {
	case 1:
		err := s.heartbeat()
		if err != nil {
			return true, xerrors.Errorf("heartbeat in response to op 1: %w", err)
		}

	// RESUME
	case 6:
		s.log.Info(s.ctx, "resumed")

		return true, nil

	// RECONNECT
	case 7:
		s.log.Info(s.ctx, "reconnect requested")

		return true, xerrors.New("reconnect")

	// INVALID_SESSION
	case 9:
		s.log.Info(s.ctx, "invalid session, reconnecting")
		s.sessID = ""
		s.persistSessID()
		s.seq = 0
		s.persistSeq()

		if s.identifyMu.IsOwner().Result == etcdserverpb.Compare_EQUAL {
			err := s.releaseIdentifyLock()
			if err != nil {
				s.log.Error(s.ctx, "release held identify lock after invalid session", slog.Error(err))
			}
		}

		return true, xerrors.New("invalid session")

	// HEARTBEAT_ACK
	case 11:
		s.lastAck = time.Now()
		return true, nil
	}

	switch ev.T {
	case "READY":
		_, sess, err := discordetf.DecodeReady(ev.D)
		if err != nil {
			return true, xerrors.Errorf("decode ready: %w", err)
		}

		s.sessID = sess
		s.log.Info(s.ctx, "ready", slog.F("sess", sess))
		s.persistSessID()

		go func() {
			time.Sleep(15 * time.Second)
			err = s.releaseIdentifyLock()
			if err != nil {
				s.log.Error(s.ctx, "release identify lock after ready", slog.Error(err))
			}
		}()

		return true, nil

	case "RESUMED":
		s.log.Info(s.ctx, "resumed")
	}

	return false, nil
}

func (s *Session) acquireIdentifyLock() error {
	err := s.identifyMu.Lock(s.ctx)
	if err != nil {
		return xerrors.Errorf("acquire identify lock: %w", err)
	}

	return nil
}

func (s *Session) releaseIdentifyLock() error {
	err := s.identifyMu.Unlock(s.ctx)
	if err != nil {
		return xerrors.Errorf("release identify lock: %w", err)
	}

	return nil
}

func (s *Session) Cancel() {
	s.cancel()
}
