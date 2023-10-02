package poller

import (
	"context"
	"log/slog"
	"time"

	event "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
)

func (p *poller) processEvents(events <-chan event.ScheduledEvent, failed chan<- event.ScheduledEvent) {
	for evt := range events {
		var err error
		switch evt.Payload.Type() {
		case event.SendNotification:
			err = p.processor.ProcessSendNotificationEvent(context.TODO(), evt)
		}
		if err != nil {
			slog.Error("failed to process event", "event_id", evt.ID.String(), "err", err.Error())

			evt.Status = event.EventStatusFailed
			evt.Attempt++
			evt.Datetime = evt.Datetime.Add(p.cfg.RetryDelay * time.Duration(evt.Attempt))

			failed <- evt
			continue
		}
		err = p.queue.Commit(context.TODO(), evt.ID)
		if err != nil {
			slog.Error("failed to commit event", "err", err.Error())
		}
	}
}

func (p *poller) processFailed(failed <-chan event.ScheduledEvent) {
	for evt := range failed {
		err := p.queue.Push(context.TODO(), evt)
		if err != nil {
			slog.Error("failed to push event to queue", "err", err.Error())
		}
	}
}

func (p *poller) pollAndSubmit(work chan<- event.ScheduledEvent) {
	events, err := p.queue.Poll(context.TODO(), p.cfg.BatchSize)
	if err != nil {
		slog.Error("queue polling error", "err", err.Error())
		return
	}
	for _, evt := range events {
		work <- evt
	}
}
