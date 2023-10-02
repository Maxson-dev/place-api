package place_repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	domain "github.com/Maxson-dev/place-api/internal/domain/place"
	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/Maxson-dev/place-api/internal/pkg/cerror"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

func (r *repo) FindOne(ctx context.Context, q database.Queryable, id int64) (domain.Place, error) {
	qb := baseSelect.
		Where(sq.Eq{"id": id})

	var dto placeDTO
	err := q.Get(ctx, &dto, qb)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Place{}, cerror.ErrNotFound
		}
		return domain.Place{}, errors.Wrap(err, sq.DebugSqlizer(qb))
	}

	return mapToPlace(dto), nil
}
