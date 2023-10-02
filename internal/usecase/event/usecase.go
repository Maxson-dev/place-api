package event

import (
	"context"

	domain "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
	"github.com/Maxson-dev/place-api/internal/infra/database"
)

type eventRepo interface {
	Save(ctx context.Context, q database.Queryable, p domain.ScheduledEvent) error
}

type usecase struct {
	db        database.PGX
	eventRepo eventRepo
}

func New(db database.PGX, eventRepo eventRepo) *usecase {
	return &usecase{
		db:        db,
		eventRepo: eventRepo,
	}
}
