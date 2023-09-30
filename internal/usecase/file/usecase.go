package fileusecase

import (
	"context"
	"io"
	"time"

	domain "github.com/Maxson-dev/place-api/internal/domain/file"
	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/google/uuid"
)

type s3Client interface {
	UploadFile(bucket, key, contentType string, body io.ReadSeeker) (string, error)
	GetSignedDownloadURL(bucket, key, name, contentType string, expire time.Duration) (string, error)
}

type fileRepo interface {
	Save(ctx context.Context, q database.Queryable, file domain.File) error
	FindOne(ctx context.Context, q database.Queryable, id uuid.UUID) (domain.File, error)
}

type Config struct {
	StorageBucket          string
	DownloadUrlLifetimeMin int64
}

type usecase struct {
	cfg      Config
	s3Client s3Client
	fileRepo fileRepo
	db       database.PGX
}

func New(db database.PGX, s3Client s3Client, fileRepo fileRepo, cfg Config) *usecase {
	return &usecase{
		cfg:      cfg,
		db:       db,
		s3Client: s3Client,
		fileRepo: fileRepo,
	}
}
