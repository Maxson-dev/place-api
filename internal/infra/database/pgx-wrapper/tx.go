package pgxwrapper

import (
	"context"

	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type txwrap struct {
	tx pgx.Tx
}

func (d *txwrap) Commit(ctx context.Context) error {
	return d.tx.Commit(ctx)
}

func (d *txwrap) Rollback(ctx context.Context) error {
	return d.tx.Rollback(ctx)
}

func (d *txwrap) Exec(ctx context.Context, sqlizer database.Sqlizer) (pgconn.CommandTag, error) {
	return exec(ctx, d.tx, sqlizer)
}

func (d *txwrap) Get(ctx context.Context, dst interface{}, sqlizer database.Sqlizer) error {
	return get(ctx, d.tx, dst, sqlizer)
}

func (d *txwrap) Select(ctx context.Context, dst interface{}, sqlizer database.Sqlizer) error {
	return selectx(ctx, d.tx, dst, sqlizer)
}

func (d *txwrap) ExecRaw(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return execRaw(ctx, d.tx, sql, arguments...)
}
