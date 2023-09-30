package fileusecase

import (
	"context"
	"io"

	domain "github.com/Maxson-dev/place-api/internal/domain/file"
)

type s3Client interface {
	UploadFile(bucket, key, contentType string, body io.ReadSeeker) (string, error)
}

type fileRepo interface {
	Save(ctx context.Context, file domain.File) error
}

type Config struct {
	StorageBucket string
}

type usecase struct {
	cfg      Config
	s3Client s3Client
	fileRepo fileRepo
}

func New(cfg Config, s3Client s3Client, fileRepo fileRepo) *usecase {
	return &usecase{
		cfg:      cfg,
		s3Client: s3Client,
		fileRepo: fileRepo,
	}
}
