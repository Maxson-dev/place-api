package poller

import (
	"context"
	"sync"

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
	BatchSize int64
	PoolSize  int64
}

type poller struct {
	stopCh  chan struct{}
	runOnce *sync.Once

	cfg       Config
	queue     queue
	processor processor
}

func New(queue queue, processor processor) *poller {
	return &poller{
		runOnce:   new(sync.Once),
		stopCh:    make(chan struct{}),
		queue:     queue,
		processor: processor,
	}
}
