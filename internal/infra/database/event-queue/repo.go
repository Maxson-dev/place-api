package eventqueue

import "github.com/Maxson-dev/place-api/internal/infra/database"

type queue struct {
	db database.PGX
}

func New(db database.PGX) *queue {
	return &queue{
		db: db,
	}
}
