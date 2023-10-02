package event_queue

import (
	"context"
	"fmt"

	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (r *repo) exists(ctx context.Context, evtID uuid.UUID) (bool, error) {
	qb := fmt.Sprintf(`select exists(select 1 from %s where id=%s);`, database.TableScheduledEvent, evtID.String())
	var exists bool
	err := r.db.Get(ctx, &exists, database.RawQuery(qb))
	if err != nil {
		return false, errors.Wrap(err, "r.db.Get")
	}
	return exists, nil
}
