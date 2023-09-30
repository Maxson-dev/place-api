package file

import (
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID        uuid.UUID
	Bucket    string
	Key       string
	Name      string
	Ext       string
	URL       string
	Size      int64
	CreatedAt time.Time
}

func New(bucket, name, key, url string, size int64) File {
	ext := filepath.Ext(name)
	return File{
		ID:        uuid.New(),
		Bucket:    bucket,
		Name:      name,
		Key:       key,
		Ext:       ext,
		URL:       url,
		Size:      size,
		CreatedAt: time.Now().UTC(),
	}
}
