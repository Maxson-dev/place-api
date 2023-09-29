package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealth health-check method
//
// Swagger spec:
// @title       Health check
// @version     1.0
// @Tags  	    api
// @host        localhost:8080
// @BasePath    /api/v1/health [get].
func (c *controller) GetHealth(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
