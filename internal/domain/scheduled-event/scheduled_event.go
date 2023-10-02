package scheduledevent

import (
	"encoding/json"
	"time"

	"github.com/Maxson-dev/place-api/internal/pkg/cerror"
	"github.com/google/uuid"
)

type ScheduledEvent struct {
	ID        uuid.UUID
	Status    Status
	Datetime  time.Time
	Payload   Payload
	CreatedAt time.Time
	Attempt   int64
}

// New scheduledEvent constructor
func New(datetime time.Time, eventType string, payload []byte) (ScheduledEvent, error) {
	switch eventType {
	case string(SendNotification):
		var p SendNotificationPayload
		err := json.Unmarshal(payload, &p)
		if err != nil {
			return ScheduledEvent{}, cerror.ErrBadInput
		}
		return ScheduledEvent{
			ID:        uuid.New(),
			Status:    EventStatusNew,
			Datetime:  datetime,
			Payload:   p,
			Attempt:   0,
			CreatedAt: time.Now().UTC(),
		}, nil
	}
	return ScheduledEvent{}, cerror.ErrBadInput
}

type Status uint8

const (
	EventStatusNew Status = iota
	EventStatusInProgress
	EventStatusDone
	EventStatusFailed
)

type Type string

const (
	SendNotification Type = "SEND_NOTIFICATION" // whatever
)

type Payload interface {
	IsScheduledEventPayload()
	Type() Type
}
