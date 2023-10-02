package poller

import (
	"time"

	event "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
)

func (p *poller) Run() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	work := make(chan event.ScheduledEvent)
	failed := make(chan event.ScheduledEvent)

	for i := int64(0); i < p.cfg.PoolSize; i++ {
		go p.processEvents(work, failed)
	}

	// release wokers after poller shutdown
	defer close(work)
	defer close(failed)

	go p.processFailed(failed)

	for {
		select {
		case <-ticker.C:
			p.pollAndSubmit(work)
		case <-p.stopCh:
			return
		}
	}
}

func (p *poller) Stop() {
	close(p.stopCh)
}
