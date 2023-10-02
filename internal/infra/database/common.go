package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var PSQL = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type PGX interface {
	Queryable
	BeginTX(ctx context.Context, opts *pgx.TxOptions) (Tx, error)
}

type Tx interface {
	Queryable
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Queryable interface {
	Exec(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error)
	Get(ctx context.Context, dst interface{}, sqlizer Sqlizer) error
	Select(ctx context.Context, dst interface{}, sqlizer Sqlizer) error
	ExecRaw(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type Sqlizer interface {
	ToSql() (sql string, args []interface{}, err error)
}

type RawQuery string

// nolint:revive
func (r RawQuery) ToSql() (string, []interface{}, error) {
	return string(r), nil, nil
}

type Execer interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}
