package pgxwrapper

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	PoolMax int64
}

type db struct {
	cfg  Config
	pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string, cfg Config) (*db, error) {
	dsn = fmt.Sprintf("%s&pool_max_conns=%d", dsn, cfg.PoolMax)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &db{cfg: cfg, pool: pool}, nil
}
