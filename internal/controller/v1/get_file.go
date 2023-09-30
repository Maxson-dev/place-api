package v1

import (
	"net/http"

	"github.com/Maxson-dev/place-api/internal/pkg/cerror"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

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
func (c *controller) GetFile(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		c.Error(ctx, err, "invalid file id", http.StatusBadRequest)
		return
	}

	file, err := c.fileUC.Get(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, cerror.ErrNotFound):
			c.Error(ctx, err, "file not found", http.StatusNotFound)
		default:
			c.Error(ctx, err, "failed to get file", http.StatusInternalServerError)
		}
		return
	}

	ctx.JSON(http.StatusOK, GetFileResponse{
		ID:   file.ID.String(),
		Name: file.Name,
		URL:  file.URL,
		Size: file.Size,
	})
}
