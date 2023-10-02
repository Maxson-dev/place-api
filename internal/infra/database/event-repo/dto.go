package event_repo

import (
	"time"

	"github.com/google/uuid"
)

var columns = []string{
	"id",
	"status",
	"type",
	"payload",
	"datetime",
	"attempt",
	"created_at",
	"updated_at",
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
