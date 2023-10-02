package eventqueue

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	event "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (q *queue) Commit(ctx context.Context, evtID uuid.UUID) error {
	qb := database.PSQL.
		Update(database.TableScheduledEvent).
		Set("status", event.ScheduledEventStatusDone).
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": evtID.String()})

	if _, err := q.db.Exec(ctx, qb); err != nil {
		return errors.Wrap(err, sq.DebugSqlizer(qb))
	}

	return nil
}
