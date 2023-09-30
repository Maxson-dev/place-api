package file

import (
	"time"

	"github.com/google/uuid"
)

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

type fileDTO struct {
	ID        uuid.UUID `db:"id"`
	Bucket    string    `db:"bucket"`
	Name      string    `db:"name"`
	S3Url     string    `db:"s3_url"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Size      int64     `db:"size"`
	IsDeleted bool      `db:"is_deleted"`
}
