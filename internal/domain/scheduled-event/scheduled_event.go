package scheduled_event

import (
	"encoding/json"
	"time"

	"github.com/Maxson-dev/place-api/internal/pkg/cerror"
	"github.com/google/uuid"
)

type ScheduledEvent struct {
	ID        uuid.UUID
	Status    ScheduledEventStatus
	Datetime  time.Time
	Payload   ScheduledEventPayload
	CreatedAt time.Time
	Attempt   int64
}

// New scheduledEvent constructor
func New(datetime time.Time, eventType string, payload []byte) (ScheduledEvent, error) {
	switch eventType {
	case string(ScheduledEventTypeSendNotification):
		var p SendNotificationPayload
		err := json.Unmarshal(payload, &p)
		if err != nil {
			return ScheduledEvent{}, cerror.ErrBadInput
		}
		return ScheduledEvent{
			ID:        uuid.New(),
			Status:    ScheduledEventStatusNew,
			Datetime:  datetime,
			Payload:   p,
			Attempt:   0,
			CreatedAt: time.Now().UTC(),
		}, nil
	}
	return ScheduledEvent{}, cerror.ErrBadInput
}

type ScheduledEventStatus uint8

const (
	ScheduledEventStatusNew ScheduledEventStatus = iota
	ScheduledEventStatusInProgress
	ScheduledEventStatusDone
	ScheduledEventStatusFailed
)

type ScheduledEventType string

const (
	ScheduledEventTypeSendNotification ScheduledEventType = "SEND_NOTIFICATION" // whatever
)

type ScheduledEventPayload interface {
	IsScheduledEventPayload()
	Type() ScheduledEventType
}
