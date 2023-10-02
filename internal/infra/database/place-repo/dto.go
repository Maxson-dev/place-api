package place_repo

import (
	"time"

	domain "github.com/Maxson-dev/place-api/internal/domain/place"
	"github.com/Maxson-dev/place-api/internal/infra/database"
)

var baseSelect = database.PSQL.
	Select(columns...).
	From(database.TablePlace)

var columns = []string{
	"id",
	"name",
	"lat",
	"lng",
	"created_at",
	"updated_at",
}

type placeDTO struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Lat       float64   `db:"lat"`
	Lng       float64   `db:"lng"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func mapToPlace(dto placeDTO) domain.Place {
	return domain.Place{
		ID: dto.ID,
		Attributes: domain.Attributes{
			Name: dto.Name,
			Coord: domain.Coord{
				Lat: dto.Lat,
				Lng: dto.Lng,
			},
			CreatedAt: dto.CreatedAt,
		},
	}
}
