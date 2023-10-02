package filerepo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	domain "github.com/Maxson-dev/place-api/internal/domain/file"
	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/pkg/errors"
)

func (r *repo) Save(ctx context.Context, q database.Queryable, file domain.File) error {
	qb := database.PSQL.
		Insert(database.TableFile).
		Columns(columns...).
		Values(
			file.ID,
			file.Size,
			file.Bucket,
			file.Key,
			file.Name,
			file.URL,
			file.CreatedAt,
			file.CreatedAt,
			false,
		)

	if _, err := q.Exec(ctx, qb); err != nil {
		return errors.Wrap(err, sq.DebugSqlizer(qb))
	}

	return nil
}
