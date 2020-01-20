package discordetf

import (
	"encoding/binary"

	"golang.org/x/xerrors"
)

type GuildCreate struct {
	Id          int64
	Guild       []byte
	MemberCount int64
	Channels    map[int64][]byte
	Emojis      map[int64][]byte
	Members     map[int64][]byte
	Presences   map[int64][]byte
	Roles       map[int64][]byte
	VoiceStates map[int64][]byte
}

func DecodeGuildCreate(buf []byte) (*GuildCreate, error) {
	var (
		d     = &decoder{buf: buf}
		gBuf  = []byte{116, 0, 0, 0, 0}
		gKeys uint32
		gc    = &GuildCreate{}
		id    int64
		err   error
	)

	id, err = d.idFromMap("id")
	if err != nil {
		return nil, xerrors.Errorf("find guild id: %w", err)
	}
	d.reset()

	err = d.checkByte(ettMap)
	if err != nil {
		return nil, xerrors.Errorf("verify map byte: %w", err)
	}

	left := d.readMapLen()
	for ; left > 0; left-- {
		start := d.off

		l, err := d.readAtomWithTag()
		if err != nil {
			return nil, xerrors.Errorf("read map key: %w", err)
		}

		key := string(d.buf[d.off-l : d.off])
		switch key {
		case "channels":
			gc.Channels, err = d.readListIntoMapByIDFixGuildID(id)
			if err != nil {
				return nil, xerrors.Errorf("read channels: %w", err)
			}

		case "emojis":
			gc.Emojis, err = d.readListIntoMapByID()
			if err != nil {
				return nil, xerrors.Errorf("read emojis: %w", err)
			}

		case "members":
			gc.Members, err = d.readListIntoMapByID()
			if err != nil {
				return nil, xerrors.Errorf("read members: %w", err)
			}

		case "presences":
			gc.Presences, err = d.readListIntoMapByID()
			if err != nil {
				return nil, xerrors.Errorf("read presences: %w", err)
			}

		case "roles":
			gc.Roles, err = d.readListIntoMapByID()
			if err != nil {
				return nil, xerrors.Errorf("read roles: %w", err)
			}

		case "voice_states":
			gc.VoiceStates, err = d.readListIntoMapByIDFixGuildID(id)
			if err != nil {
				return nil, xerrors.Errorf("read voice states: %w", err)
			}

		case "id":
			gc.Id, err = d.readSmallBigWithTagToInt64()
			if err != nil {
				return nil, xerrors.Errorf("read id: %w", err)
			}

			gBuf = append(gBuf, d.buf[start:d.off]...)
			gKeys++

		case "member_count":
			gc.MemberCount, err = d.readInteger()
			if err != nil {
				return nil, xerrors.Errorf("read member_count: %w", err)
			}

			gBuf = append(gBuf, d.buf[start:d.off]...)
			gKeys++

		default:
			err := d.readTerm()
			if err != nil {
				return nil, err
			}

			gBuf = append(gBuf, d.buf[start:d.off]...)
			gKeys++
		}
	}

	// fix length
	binary.BigEndian.PutUint32(gBuf[1:5], gKeys)
	gc.Guild = gBuf

	return gc, nil
}

type GuildBan struct {
	User  int64
	Guild int64
	Raw   []byte
}

func DecodeGuildBan(buf []byte) (*GuildBan, error) {
	var (
		gb  = &GuildBan{}
		d   = &decoder{buf: buf}
		err error
	)

	gb.User, gb.Raw, err = d.readMapWithIDIntoSlice()
	if err != nil {
		return nil, xerrors.Errorf("read user id: %w", err)
	}

	d.reset()
	gb.Guild, err = d.idFromMap("guild_id")
	if err != nil {
		return nil, xerrors.Errorf("read guild id: %w", err)
	}

	return gb, err
}
