package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostFileResponse struct {
	IDs []string `json:"ids"`
}

// PostFile
//
// @Summary     Upload file
// @Description Method for uploading file to storage
// @ID          PostFile
// @Tags  	    file_api
// @Accept      mpfd
// @Produce     json
// @Param       file formData file true "Body with file"
// @Success     200 {object} PostFileResponse
// @Failure     400 {object} Error
// @Failure     500 {object} Error
// @Router      /api/v1/file [post]
func (c *controller) PostFile(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		c.Error(ctx, "invalid payload", http.StatusBadRequest)
		return
	}

	files := form.File["upload[]"]

	fileIDs, err := c.fileUC.Upload(ctx, files)
	if err != nil {
		c.Error(ctx, "failed to upload file", http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, PostFileResponse{IDs: fileIDs})
}
