package v1

import (
	"context"
	"mime/multipart"

	domain "github.com/Maxson-dev/place-api/internal/domain/file"
	"github.com/google/uuid"
)

type fileUC interface {
	Get(ctx context.Context, id uuid.UUID) (domain.File, error)
	Upload(ctx context.Context, files []*multipart.FileHeader) ([]string, error)
}

type placeUC interface {
	Create(ctx context.Context, name string, lat, lng float64) (int64, error)
}

type controller struct {
	fileUC  fileUC
	placeUC placeUC
}

func New(fileUC fileUC, placeUC placeUC) *controller {
	return &controller{
		fileUC:  fileUC,
		placeUC: placeUC,
	}
}
