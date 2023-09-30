package v1

import (
	"context"
	"mime/multipart"
)

type fileUC interface {
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
