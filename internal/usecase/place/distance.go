package place

import (
	"context"

	"github.com/Maxson-dev/place-api/internal/domain/place"
	"github.com/pkg/errors"
)

func (u *usecase) GetDistance(ctx context.Context, id int64, lat float64, lng float64) (mi, km float64, err error) {
	obj, err := u.placeRepo.FindOne(ctx, u.db, id)
	if err != nil {
		return 0, 0, errors.Wrap(err, "placeRepo.FindOne")
	}

	mi, km = place.Distance(obj.Coord, place.Coord{Lat: lat, Lng: lng})

	return mi, km, nil
}
