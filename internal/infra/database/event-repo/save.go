package event_repo

import (
	"context"
	"encoding/json"

	sq "github.com/Masterminds/squirrel"
	domain "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
)

func (r *repo) Save(ctx context.Context, q database.Queryable, evt domain.ScheduledEvent) error {
	return r.save(ctx, q, evt)
}

func (r *repo) save(ctx context.Context, q database.Queryable, evt domain.ScheduledEvent) error {
	pld, err := json.Marshal(evt.Payload)
	if err != nil {
		return errors.Wrap(err, "could not marshal payload")
	}

	qb := database.PSQL.
		Insert(database.TableScheduledEvent).
		Columns(
			"id",
			"status",
			"type",
			"payload",
			"datetime",
			"attempt",
			"created_at",
			"updated_at",
		).
		Values(
			evt.ID,
			evt.Status,
			string(evt.Payload.Type()),
			&pgtype.JSONB{Bytes: pld, Status: pgtype.Present},
			evt.Datetime,
			evt.Attempt,
			evt.CreatedAt,
			evt.CreatedAt,
		)

	if _, err := q.Exec(ctx, qb); err != nil {
		return errors.Wrap(err, sq.DebugSqlizer(qb))
	}

	return nil
}
