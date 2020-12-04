package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/xerrors"
)

func (s *Server) getGuildRole(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	guild, err := guildParam(p)
	if err != nil {
		return xerrors.Errorf("read guild param: %w", err)
	}
	ro, err := s.db.GetGuildRole(r.Context(), guild, roleParam(p))
	if err != nil {
		return xerrors.Errorf("read role: %w", err)
	}

	if ro == nil {
		return ErrNotFound
	}

	return s.writeTerm(w, ro)
}

func (s *Server) getGuildRoles(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	if len(r.URL.Query()["id"]) > 0 {
		return s.getGuildRoleSlice(w, r, p)
	}
	guild, err := guildParam(p)
	if err != nil {
		return xerrors.Errorf("read guild param: %w", err)
	}
	ros, err := s.db.GetGuildRoles(r.Context(), guild)
	if err != nil {
		return xerrors.Errorf("read roles: %w", err)
	}

	return s.writeTerms(w, ros)
}

func (s *Server) getGuildRoleSlice(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	g, err := guildParam(p)
	if err != nil {
		return xerrors.Errorf("read guild param: %w", err)
	}
	var (
		ctx = r.Context()
		rs  = r.URL.Query()["id"]
		ros = make([][]byte, 0, len(rs))
	)

	for _, e := range rs {
		rr, err := strconv.ParseInt(e, 10, 64)
		if err != nil {
			return xerrors.Errorf("parse role id: %w", err)
		}

		rol, err := s.db.GetGuildRole(ctx, g, rr)
		if err != nil {
			if xerrors.Is(err, ErrNotFound) {
				rol, _ = json.Marshal(EmptyObj{Id: e, IsEmpty: true})
			} else {
				return xerrors.Errorf("get role: %w", err)
			}
		}

		ros = append(ros, rol)
	}

	return s.writeTerms(w, ros)
}

func roleParam(p httprouter.Params) int64 {
	r := p.ByName("role")
	ri, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	return ri
}
