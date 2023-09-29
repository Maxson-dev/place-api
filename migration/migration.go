package migration

import (
	"context"
	"database/sql"
	"embed"
	"log/slog"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

func Migrate(ctx context.Context, uri string) error {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return err
	}
	defer func() {
		err := db.Close()
		if err != nil {
			slog.ErrorContext(ctx, "could not close db connection: %s", err)
		}
	}()

	goose.SetBaseFS(migrations)
	err = goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	if errUp := goose.Up(db, "migrations"); errUp != nil {
		return errUp
	}

	return nil
}
