package pgxwrapper

import (
	"context"

	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (d *dbwrap) BeginTX(ctx context.Context, opts *pgx.TxOptions) (database.Tx, error) {
	tx, err := d.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return &txwrap{tx: tx}, nil
}

func (d *dbwrap) Exec(ctx context.Context, sqlizer database.Sqlizer) (pgconn.CommandTag, error) {
	return exec(ctx, d.pool, sqlizer)
}

func (d *dbwrap) Get(ctx context.Context, dst interface{}, sqlizer database.Sqlizer) error {
	return get(ctx, d.pool, dst, sqlizer)
}

func (d *dbwrap) Select(ctx context.Context, dst interface{}, sqlizer database.Sqlizer) error {
	return selectx(ctx, d.pool, dst, sqlizer)
}

func (d *dbwrap) ExecRaw(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return execRaw(ctx, d.pool, sql, arguments...)
}

func execRaw(ctx context.Context, ex database.Execer, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	return ex.Exec(ctx, sql, arguments...)
}

func exec(ctx context.Context, ex database.Execer, sqlizer database.Sqlizer) (pgconn.CommandTag, error) {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return ex.Exec(ctx, sql, args...)
}

func get(ctx context.Context, q pgxscan.Querier, dst interface{}, sqlizer database.Sqlizer) error {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	return pgxscan.Get(ctx, q, dst, sql, args...)
}

func selectx(ctx context.Context, q pgxscan.Querier, dst interface{}, sqlizer database.Sqlizer) error {
	sql, args, err := sqlizer.ToSql()
	if err != nil {
		return err
	}

	return pgxscan.Select(ctx, q, dst, sql, args...)
}
