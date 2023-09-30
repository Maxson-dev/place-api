package v1

import "github.com/gin-gonic/gin"

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (c *controller) Error(ctx *gin.Context, msg string, statusCode int) {
	ctx.JSON(statusCode, Error{
		Status:  "error",
		Message: msg,
	})
}
