package v1

import "github.com/gin-gonic/gin"

type PostPlaceRequest struct {
	Lat int64 `json:"lat"`
	Lng int64 `json:"lng"`
}

type PostPlaceResponse struct {
	ID string `json:"id"`
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
func (c *controller) PostPlace(ctx *gin.Context) {}
