package eventqueue

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	event "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/spaolacci/murmur3"
)

const pollEventsLock = "poll_scheduled_events_lock"

func (q *queue) Poll(ctx context.Context, max int64) ([]event.ScheduledEvent, error) {
	tx, err := q.db.BeginTX(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not begin transaction")
	}
	defer database.RollbackTx(ctx, tx)

	lockq := fmt.Sprintf("select pg_advisory_xact_lock(%d)", murmur3.Sum32([]byte(pollEventsLock)))

	// block other service instances from receiving from queue
	_, err = tx.ExecRaw(ctx, lockq)
	if err != nil {
		return nil, errors.Wrap(err, "could not take pg_advisory_xact_lock")
	}

	clms := strings.Join(columns, ",")

	qb := fmt.Sprintf(
		`
	with batch as (
		select %s from %s
    	where status = %d -- new
    	limit %d
	) 
	update scheduled_event
		set status = %d -- in progress
	from batch
	where scheduled_event.id = batch.id
	returning %s;`,
		clms,
		database.TableScheduledEvent,
		event.EventStatusNew,
		max,
		event.EventStatusInProgress,
		clms,
	)

	var dtos []scheduledEventDTO
	err = q.db.Select(ctx, &dtos, database.RawQuery(qb))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []event.ScheduledEvent{}, nil
		}
		return nil, errors.Wrap(err, qb)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, errors.Wrap(err, "could not commit transaction")
	}

	return mapToScheduledEvents(dtos)
}

func mapToScheduledEvents(dtos []scheduledEventDTO) ([]event.ScheduledEvent, error) {
	events := make([]event.ScheduledEvent, 0, len(dtos))
	for _, dto := range dtos {
		evt, err := mapToScheduledEvent(dto)
		if err != nil {
			return nil, err
		}
		events = append(events, evt)
	}
	return events, nil
}

func mapToScheduledEvent(dto scheduledEventDTO) (event.ScheduledEvent, error) {
	var err error
	var pld event.Payload
	switch event.Type(dto.Type) {
	case event.SendNotification:
		var p event.SendNotificationPayload
		err = json.Unmarshal(dto.Payload, &p)
		pld = p
	}
	if err != nil {
		return event.ScheduledEvent{}, errors.Wrap(err, "could not unmarshal payload")
	}

	return event.ScheduledEvent{
		ID:        dto.ID,
		Status:    event.Status(dto.Status),
		Payload:   pld,
		Datetime:  dto.Datetime,
		Attempt:   dto.Attempt,
		CreatedAt: dto.CreatedAt,
	}, nil
}
