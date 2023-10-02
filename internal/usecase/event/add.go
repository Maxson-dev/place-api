package event

import (
	"context"
	"time"

	event "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
	"github.com/pkg/errors"
)

func (u *usecase) Add(ctx context.Context, eventType string, datetime time.Time, payload []byte) (event.ScheduledEvent, error) {
	evt, err := event.New(datetime, eventType, payload)
	if err != nil {
		return event.ScheduledEvent{}, errors.Wrap(err, "event.New")
	}
	err = u.eventRepo.Save(ctx, u.db, evt)
	if err != nil {
		return event.ScheduledEvent{}, errors.Wrap(err, "u.eventRepo.Save")
	}
	return evt, nil
}
