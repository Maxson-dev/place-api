package file_repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	domain "github.com/Maxson-dev/place-api/internal/domain/file"
	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/Maxson-dev/place-api/internal/pkg/cerror"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *repo) FindOne(ctx context.Context, q database.Queryable, id uuid.UUID) (domain.File, error) {
	qb := baseSelect.
		Where(sq.Eq{"id": id})

	var dto fileDTO
	err := q.Get(ctx, &dto, qb)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.File{}, cerror.ErrNotFound
		}
		return domain.File{}, errors.Wrap(err, sq.DebugSqlizer(qb))
	}

	return mapToFile(dto), nil
}
