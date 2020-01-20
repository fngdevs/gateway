package discordetf

import (
	"golang.org/x/xerrors"
)

type Channel struct {
	Id    int64
	Guild int64
	Raw   []byte
}

func DecodeChannel(buf []byte) (*Channel, error) {
	var (
		ch  = &Channel{}
		d   = &decoder{buf: buf}
		err error
	)

	ch.Id, ch.Raw, err = d.readMapWithIDIntoSlice()
	if err != nil {
		return nil, xerrors.Errorf("extract id: %w", err)
	}

	d.reset()
	ch.Guild, err = d.guildIDFromMap()
	if err != nil {
		return nil, xerrors.Errorf("extract guild_id: %w", err)
	}

	return ch, err
}

type VoiceState struct {
	User  int64
	Guild int64
	Raw   []byte
}

func DecodeVoiceState(buf []byte) (*VoiceState, error) {
	var (
		vs  = &VoiceState{}
		d   = &decoder{buf: buf}
		err error
	)

	vs.User, vs.Raw, err = d.readMapWithIDIntoSlice()
	if err != nil {
		return nil, xerrors.Errorf("extract user_id: %w", err)
	}

	d.reset()
	vs.Guild, err = d.idFromMap("guild_id")
	if err != nil {
		return nil, xerrors.Errorf("extract guild_id: %w", err)
	}

	return vs, err
}
