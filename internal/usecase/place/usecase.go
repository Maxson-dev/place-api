package place

import (
	"context"

	"github.com/Maxson-dev/place-api/internal/domain/place"
	"github.com/Maxson-dev/place-api/internal/infra/database"
)

type placeRepo interface {
	Save(ctx context.Context, q database.Queryable, p place.Attributes) (int64, error)
	FindOne(ctx context.Context, q database.Queryable, id int64) (place.Place, error)
}

type usecase struct {
	db        database.PGX
	placeRepo placeRepo
}

func New(db database.PGX, placeRepo placeRepo) *usecase {
	return &usecase{
		db:        db,
		placeRepo: placeRepo,
	}
}
