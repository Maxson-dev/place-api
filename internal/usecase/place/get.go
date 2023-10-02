package place

import (
	"context"

	domain "github.com/Maxson-dev/place-api/internal/domain/place"
	"github.com/pkg/errors"
)

func (u *usecase) Get(ctx context.Context, id int64) (domain.Place, error) {
	place, err := u.placeRepo.FindOne(ctx, u.db, id)
	if err != nil {
		return domain.Place{}, errors.Wrap(err, "placeRepo.FindOne")
	}
	return place, nil
}
