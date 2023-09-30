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

type controller struct {
	fileUC fileUC
}

func New(fileUC fileUC) *controller {
	return &controller{
		fileUC: fileUC,
	}
}
