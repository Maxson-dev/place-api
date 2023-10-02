package eventqueue

import (
	"time"

	"github.com/google/uuid"
)

var columns = []string{
	"scheduled_event.id",
	"scheduled_event.status",
	"scheduled_event.type",
	"scheduled_event.payload",
	"scheduled_event.datetime",
	"scheduled_event.attempt",
	"scheduled_event.created_at",
	"scheduled_event.updated_at",
}

type scheduledEventDTO struct {
	ID        uuid.UUID `db:"id"`
	Status    int64     `db:"status"`
	Attempt   int64     `db:"attempt"`
	Type      string    `db:"type"`
	Payload   []byte    `db:"payload"`
	Datetime  time.Time `db:"datetime"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
