package fileusecase

import (
	"context"
	"io"

	domain "github.com/Maxson-dev/place-api/internal/domain/file"
	"github.com/Maxson-dev/place-api/internal/infra/database"
)

type s3Client interface {
	UploadFile(bucket, key, contentType string, body io.ReadSeeker) (string, error)
}

type fileRepo interface {
	Save(ctx context.Context, q database.Queryable, file domain.File) error
}

type usecase struct {
	storageBucket string
	s3Client      s3Client
	fileRepo      fileRepo
	db            database.PGX
}

func New(storageBucket string, s3Client s3Client, fileRepo fileRepo) *usecase {
	return &usecase{
		storageBucket: storageBucket,
		s3Client:      s3Client,
		fileRepo:      fileRepo,
	}
}
