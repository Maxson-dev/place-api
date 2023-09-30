package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostPlaceRequest struct {
	Name string  `json:"name" binding:"required"`
	Lat  float64 `json:"lat" binding:"required"`
	Lng  float64 `json:"lng" binding:"required"`
}

type PostPlaceResponse struct {
	ID int64 `json:"id"`
}

// PostPlace
//
// @Summary     Create object
// @Description Method for object creation
// @ID          PostPlace
// @Tags  	    place_api
// @Accept      json
// @Produce     json
// @Param       request body PostPlaceRequest true "Create object"
// @Success     200 {object} PostPlaceResponse
// @Failure     400 {object} Error
// @Failure     500 {object} Error
// @Router      /api/v1/place [post]
func (c *controller) PostPlace(ctx *gin.Context) {
	var req PostPlaceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.Error(ctx, err, "invalid object", http.StatusBadRequest)
		return
	}

	id, err := c.placeUC.Create(ctx, req.Name, req.Lat, req.Lng)
	if err != nil {
		c.Error(ctx, err, "failed to create object", http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, PostPlaceResponse{ID: id})
}
