package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type v1API interface {
	PostFile(ctx *gin.Context)
	GetFile(ctx *gin.Context)
}

type HTTPConfig struct {
	Port                int64
	MaxMultipartSizeMib int64
}

type service struct {
	engine *gin.Engine
	cfg    HTTPConfig
}

// New creates a service
// Swagger spec:
// @title       Place-API Service.
// @version     0.1.0
// @host        localhost:8080
// @BasePath  	/api/v1
func New(engine *gin.Engine, v1 v1API, cfg HTTPConfig) *service {
	srv := &service{
		cfg:    cfg,
		engine: engine,
	}

	api := srv.engine.Group("/api/")
	{
		apiV1 := api.Group("/v1/")
		{
			{
				apiV1.POST("/file", v1.PostFile)
				apiV1.GET("/file/:id", v1.GetFile)
			}
			{
				apiV1.POST("/place")
				apiV1.GET("/place/:id")
				apiV1.GET("/place/:id/distance")
			}
			{
				apiV1.POST("/event")
			}
			{
				apiV1.GET("/health")
			}
		}

	}

	return srv
}

func (s *service) Run() error {
	return s.engine.Run(fmt.Sprintf(":%d", s.cfg.Port))
}
