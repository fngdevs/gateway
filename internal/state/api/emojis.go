package api

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/xerrors"
)

func (s *Server) getGuildEmoji(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	guild, err := guildParam(p)
	if err != nil {
		return xerrors.Errorf("read guild param: %w", err)
	}
	c, err := s.db.GetGuildEmoji(r.Context(), guild, emojiParam(p))
	if err != nil {
		return xerrors.Errorf("read emoji: %w", err)
	}

	if c == nil {
		return ErrNotFound
	}

	return s.writeTerm(w, c)
}

func emojiParam(p httprouter.Params) int64 {
	e := p.ByName("emoji")
	eid, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	return eid
}

func (s *Server) getGuildEmojis(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	guild, err := guildParam(p)
	if err != nil {
		return xerrors.Errorf("read guild param: %w", err)
	}
	cs, err := s.db.GetGuildEmojis(r.Context(), guild)
	if err != nil {
		return xerrors.Errorf("read guild emojis: %w", err)
	}

	return s.writeTerms(w, cs)
}
