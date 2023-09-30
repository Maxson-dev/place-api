package v1

import "github.com/gin-gonic/gin"

type GetFileResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
	Size int64  `json:"size"`
}

// GetFile
//
// @Summary     Get file info
// @Description Method returns file meta and temporary url to download
// @ID          GetFile
// @Tags  	    file_api
// @Produce     json
// @Param       id path string true "File ID"
// @Success     200 {object} GetFileResponse
// @Failure     400 {object} Error
// @Failure     404 {object} Error
// @Failure     500 {object} Error
// @Router      /api/v1/file/{id} [get]
func (c *controller) GetFile(ctx *gin.Context) {}
