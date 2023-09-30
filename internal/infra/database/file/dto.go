package file

import (
	"path/filepath"
	"time"

	domain "github.com/Maxson-dev/place-api/internal/domain/file"
	"github.com/Maxson-dev/place-api/internal/infra/database"
	"github.com/google/uuid"
)

var baseSelect = database.PSQL.
	Select(columns...).
	From(database.TableFile)

var columns = []string{
	"id",
	"size",
	"bucket",
	"key",
	"name",
	"s3_url",
	"created_at",
	"updated_at",
	"is_deleted",
}

// nolint:unused
type fileDTO struct {
	ID        uuid.UUID `db:"id"`
	Bucket    string    `db:"bucket"`
	Key       string    `db:"key"`
	Name      string    `db:"name"`
	S3Url     string    `db:"s3_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Size      int64     `db:"size"`
	IsDeleted bool      `db:"is_deleted"`
}

func mapToFile(dto fileDTO) domain.File {
	return domain.File{
		ID:        dto.ID,
		Bucket:    dto.Bucket,
		Key:       dto.Key,
		Name:      dto.Name,
		Ext:       filepath.Ext(dto.Name),
		URL:       dto.S3Url,
		Size:      dto.Size,
		CreatedAt: dto.CreatedAt,
	}
}
