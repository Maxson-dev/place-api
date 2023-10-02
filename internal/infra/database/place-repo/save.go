package place_repo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/Maxson-dev/place-api/internal/domain/place"
	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/pkg/errors"
)

func (r *repo) Save(ctx context.Context, q database.Queryable, p place.Attributes) (int64, error) {
	qb := database.PSQL.
		Insert(database.TablePlace).
		Columns(
			"name",
			"lat",
			"lng",
			"created_at",
			"updated_at",
		).
		Values(
			p.Name,
			p.Coord.Lat,
			p.Coord.Lng,
			p.CreatedAt,
			p.CreatedAt,
		).Suffix("RETURNING id")

	var id int64
	if err := q.Get(ctx, &id, qb); err != nil {
		return 0, errors.Wrap(err, sq.DebugSqlizer(qb))
	}

	return id, nil
}
