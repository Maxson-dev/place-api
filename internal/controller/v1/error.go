package v1

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (c *controller) Error(ctx *gin.Context, err error, msg string, statusCode int) {
	if err != nil {
		slog.Error(fmt.Sprintf("http handler error: %s", err.Error()))
	}
	ctx.JSON(statusCode, Error{
		Status:  "error",
		Message: msg,
	})
}
