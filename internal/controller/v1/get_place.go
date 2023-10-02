package v1

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/Maxson-dev/place-api/internal/pkg/cerror"
	"github.com/Maxson-dev/place-api/internal/pkg/serd"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type GetPlaceResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	CreatedAt time.Time `json:"created_at"`
}

// GetPlace
//
// @Summary     Get place
// @Description Method returns place by id
// @ID          GetPlace
// @Tags  	    place_api
// @Produce     json
// @Param       id path string true "Place ID"
// @Success     200 {object} GetPlaceResponse
// @Failure     400 {object} Error
// @Failure     404 {object} Error
// @Failure     500 {object} Error
// @Router      /api/v1/place/{id} [get]
func (c *controller) GetPlace(ctx *gin.Context) {
	idparam := ctx.Param("id")
	id, err := serd.ParseIntID(idparam)
	if err != nil {
		c.Error(ctx, nil, "invalid place id", http.StatusBadRequest)
		return
	}

	slog.Debug("[GetPlace] idparam: ", idparam)

	place, err := c.placeUC.Get(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, cerror.ErrNotFound):
			c.Error(ctx, err, "place not found", http.StatusNotFound)
		default:
			c.Error(ctx, err, "failed to get place", http.StatusInternalServerError)
		}
		return
	}

	ctx.JSON(http.StatusOK, GetPlaceResponse{
		ID:        place.ID,
		Name:      place.Name,
		Lat:       place.Lat,
		Lng:       place.Lng,
		CreatedAt: place.CreatedAt,
	})
}
