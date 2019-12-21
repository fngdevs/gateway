package statefdb

import (
	"context"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

func (db *DB) SetGuildMembers(ctx context.Context, guild int64, raws map[int64][]byte) error {
	return db.setGuildETFs(guild, raws, db.fmtGuildMemberKey)
}

func (db *DB) DeleteGuildMembers(ctx context.Context, guild int64) error {
	pre, _ := fdb.PrefixRange(db.fmtGuildMemberPrefix(guild))

	return db.Transact(func(t fdb.Transaction) error {
		t.ClearRange(pre)
		return nil
	})
}

func (db *DB) SetGuildMember(ctx context.Context, guild, user int64, raw []byte) error {
	return db.Transact(func(t fdb.Transaction) error {
		t.Set(db.fmtGuildMemberKey(guild, user), raw)
		return nil
	})
}

func (db *DB) GetGuildMember(ctx context.Context, guild, user int64) ([]byte, error) {
	var m []byte

	err := db.Transact(func(t fdb.Transaction) error {
		m = t.Get(db.fmtGuildMemberKey(guild, user)).MustGet()
		return nil
	})
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (db *DB) GetGuildMembers(ctx context.Context, guild int64) ([][]byte, error) {
	var (
		raws   []fdb.KeyValue
		out    [][]byte
		pre, _ = fdb.PrefixRange(db.fmtGuildMemberPrefix(guild))
	)

	err := db.ReadTransact(func(t fdb.ReadTransaction) error {
		raws = t.Snapshot().GetRange(pre, FDBRangeWantAll).GetSliceOrPanic()
		return nil
	})
	if err != nil {
		return nil, err
	}

	out = make([][]byte, len(raws))
	for i, e := range raws {
		out[i] = e.Value
	}

	return out, err
}

func (db *DB) DeleteGuildMember(ctx context.Context, guild, user int64) error {
	return db.Transact(func(t fdb.Transaction) error {
		t.Clear(db.fmtGuildMemberKey(guild, user))
		return nil
	})
}
