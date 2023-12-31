package eventqueue

import (
	"context"
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"
	event "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
)

func (q *queue) Push(ctx context.Context, evt event.ScheduledEvent) error {
	exists, err := q.exists(ctx, evt.ID)
	if err != nil {
		return errors.Wrap(err, "r.exists")
	}
	if !exists {
		return q.save(ctx, q.db, evt)
	}

	pld, err := json.Marshal(evt.Payload)
	if err != nil {
		return errors.Wrap(err, "could not marshal payload")
	}

	qb := database.PSQL.
		Update(database.TableScheduledEvent).
		Set("status", evt.Status).
		Set("attempt", evt.Attempt).
		Set("datetime", evt.Datetime).
		Set("payload", &pgtype.JSONB{Bytes: pld, Status: pgtype.Present}).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": evt.ID.String()})

	if _, err := q.db.Exec(ctx, qb); err != nil {
		return errors.Wrap(err, sq.DebugSqlizer(qb))
	}

	return nil
}
