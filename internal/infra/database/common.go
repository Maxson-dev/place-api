package database

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
)

var PSQL = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Queryable interface {
	Exec(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error)
	Get(ctx context.Context, dst interface{}, sqlizer Sqlizer) error
	Select(ctx context.Context, dst interface{}, sqlizer Sqlizer) error
	ExecRaw(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type Sqlizer interface {
	ToSql() (sql string, args []interface{}, err error)
}
