package fileusecase

import (
	"context"
	"mime/multipart"

	"github.com/Maxson-dev/place-api/internal/domain/file"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
)

const octetStream = "application/octet-stream"

func (u *usecase) Upload(ctx context.Context, files []*multipart.FileHeader) ([]string, error) {
	res := make([]string, 0, len(files))

	for _, f := range files {
		id, err := u.upload(ctx, f)
		if err != nil {
			return nil, errors.Wrap(err, "u.upload")
		}
		res = append(res, id)
	}

	return res, nil
}

func (u *usecase) upload(ctx context.Context, f *multipart.FileHeader) (string, error) {
	if f == nil {
		return "", nil
	}

	body, err := f.Open()
	if err != nil {
		return "", errors.Wrap(err, "f.Open")
	}

	url, err := u.s3Client.UploadFile(u.storageBucket, f.Filename, octetStream, body)
	if err != nil {
		return "", errors.Wrap(err, "u.s3Client.UploadFile")
	}

	header := file.New(u.storageBucket, f.Filename, ulid.Make().String(), url, f.Size)

	err = u.fileRepo.Save(ctx, u.db, header)
	if err != nil {
		return "", errors.Wrap(err, "u.fileRepo.Save")
	}

	return header.ID.String(), nil
}
