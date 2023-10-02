package v1

import (
	"net/http"
	"strconv"

	"github.com/Maxson-dev/place-api/internal/pkg/cerror"
	"github.com/Maxson-dev/place-api/internal/pkg/serd"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type GetDistanceResponse struct {
	Km float64 `json:"km"`
	Mi float64 `json:"mi"`
}

// GetDistance
//
// @Summary     Get distance
// @Description Method returns distance between object and point
// @ID          GetDistance
// @Tags  	    place_api
// @Produce     json
// @Param       id path string true "Place ID"
// @Param       lat query string true "latitude"
// @Param       lng query string true "longitude"
// @Success     200 {object} GetDistanceResponse
// @Failure     400 {object} Error
// @Failure     404 {object} Error
// @Failure     500 {object} Error
// @Router      /place/{id}/distance [get]
func (c *controller) GetDistance(ctx *gin.Context) {
	lat, err := strconv.ParseFloat(ctx.Query("lat"), 64)
	if err != nil {
		c.Error(ctx, err, "invalid latitude", http.StatusBadRequest)
		return
	}

	lng, err := strconv.ParseFloat(ctx.Query("lng"), 64)
	if err != nil {
		c.Error(ctx, err, "invalid longitude", http.StatusBadRequest)
		return
	}

	id, err := serd.ParseIntID(ctx.Param("id"))
	if err != nil {
		c.Error(ctx, nil, "invalid object id", http.StatusBadRequest)
		return
	}

	mi, km, err := c.placeUC.GetDistance(ctx, id, lat, lng)
	if err != nil {
		switch {
		case errors.Is(err, cerror.ErrNotFound):
			c.Error(ctx, err, "object not found", http.StatusNotFound)
		default:
			c.Error(ctx, err, "failed to calculate distance", http.StatusInternalServerError)
		}
		return
	}

	ctx.JSON(http.StatusOK, GetDistanceResponse{
		Km: km,
		Mi: mi,
	})
}
