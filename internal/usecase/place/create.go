package place

import (
	"context"

	domain "github.com/Maxson-dev/place-api/internal/domain/place"
	"github.com/pkg/errors"
)

func (u *usecase) Create(ctx context.Context, name string, lat, lng float64) (int64, error) {
	attrs := domain.NewAttributes(name, lat, lng)

	id, err := u.placeRepo.Save(ctx, u.db, attrs)
	if err != nil {
		return 0, errors.Wrap(err, "placeRepo.Save")
	}

	return id, nil
}
