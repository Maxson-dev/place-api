package poller

import (
	"context"
	"time"

	event "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
	"github.com/google/uuid"
)

type queue interface {
	Poll(ctx context.Context, max int64) ([]event.ScheduledEvent, error)
	Commit(ctx context.Context, evtID uuid.UUID) error

	Push(ctx context.Context, evt event.ScheduledEvent) error
}

type processor interface {
	ProcessSendNotificationEvent(ctx context.Context, evt event.ScheduledEvent) error
}

type Config struct {
	BatchSize  int64
	PoolSize   int64
	RetryDelay time.Duration
}

type poller struct {
	stopCh chan struct{}

	cfg       Config
	queue     queue
	processor processor
}

func New(queue queue, processor processor, cfg Config) *poller {
	return &poller{
		cfg:       cfg,
		stopCh:    make(chan struct{}),
		queue:     queue,
		processor: processor,
	}
}
