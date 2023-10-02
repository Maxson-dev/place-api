package v1

import (
	"context"
	"mime/multipart"
	"time"

	filedomain "github.com/Maxson-dev/place-api/internal/domain/file"
	placedomain "github.com/Maxson-dev/place-api/internal/domain/place"
	eventdomain "github.com/Maxson-dev/place-api/internal/domain/scheduled-event"
	"github.com/google/uuid"
)

type fileUC interface {
	Get(ctx context.Context, id uuid.UUID) (filedomain.File, error)
	Upload(ctx context.Context, files []*multipart.FileHeader) ([]string, error)
}

type placeUC interface {
	Get(ctx context.Context, id int64) (placedomain.Place, error)
	Create(ctx context.Context, name string, lat, lng float64) (int64, error)
	GetDistance(ctx context.Context, id int64, lat float64, lng float64) (mi, km float64, err error)
}

type eventUC interface {
	Add(ctx context.Context, eventType string, datetime time.Time, payload []byte) (eventdomain.ScheduledEvent, error)
}

type controller struct {
	fileUC  fileUC
	placeUC placeUC
	eventUC eventUC
}

func New(fileUC fileUC, placeUC placeUC, eventUC eventUC) *controller {
	return &controller{
		fileUC:  fileUC,
		placeUC: placeUC,
		eventUC: eventUC,
	}
}
