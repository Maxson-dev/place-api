package event_queue

import "github.com/Maxson-dev/place-api/internal/infra/database"

type repo struct {
	db database.PGX
}

func New(db database.PGX) *repo {
	return &repo{
		db: db,
	}
}
