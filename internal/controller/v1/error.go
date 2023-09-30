package v1

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (c *controller) Error(ctx *gin.Context, err error, msg string, statusCode int) {
	slog.Error("http handler error: %s", err.Error())
	ctx.JSON(statusCode, Error{
		Status:  "error",
		Message: msg,
	})
}
