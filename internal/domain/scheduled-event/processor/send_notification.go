package processor

import (
	"context"
	"log/slog"
	"time"

	event "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
)

func (p *processor) ProcessSendNotificationEvent(ctx context.Context, evt event.ScheduledEvent) error {
	// do some io-bound work
	time.Sleep(time.Second)

	slog.Info("send notification event processed", "event_id", evt.ID)

	return nil
}
