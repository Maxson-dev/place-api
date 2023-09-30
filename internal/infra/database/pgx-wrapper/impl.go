package pgxwrapper

import (
	"context"

	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgconn"
)

func (d *db) Exec(ctx context.Context, sqlizer database.Sqlizer) (pgconn.CommandTag, error) {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return d.pool.Exec(ctx, sql, args...)
}

func (d *db) Get(ctx context.Context, dst interface{}, sqlizer database.Sqlizer) error {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	return pgxscan.Get(ctx, d.pool, dst, sql, args...)
}

func (d *db) Select(ctx context.Context, dst interface{}, sqlizer database.Sqlizer) error {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	return pgxscan.Select(ctx, d.pool, dst, sql, args...)
}

func (d *db) ExecRaw(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return d.pool.Exec(ctx, sql, arguments...)
}
