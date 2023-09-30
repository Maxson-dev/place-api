package fileusecase

import (
	"context"
	"time"

	domain "github.com/Maxson-dev/place-api/internal/domain/file"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (u *usecase) Get(ctx context.Context, id uuid.UUID) (domain.File, error) {
	file, err := u.fileRepo.FindOne(ctx, u.db, id)
	if err != nil {
		return domain.File{}, errors.Wrap(err, "u.fileRepo.FindOne")
	}

	url, err := u.s3Client.GetSignedDownloadURL(
		file.Bucket, file.Key, file.Name,
		octetStream,
		time.Duration(u.cfg.DownloadURLLifetimeMin)*time.Minute,
	)
	if err != nil {
		return domain.File{}, errors.Wrap(err, "u.s3Client.GetSignedDownloadURL")
	}

	file.URL = url

	return file, nil
}
