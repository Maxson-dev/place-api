package controller

import (
	"fmt"

	_ "github.com/Maxson-dev/place-api/api" // swagger docs
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type v1API interface {
	PostFile(ctx *gin.Context)
	GetFile(ctx *gin.Context)
	PostPlace(ctx *gin.Context)
	GetPlace(ctx *gin.Context)
	GetDistance(ctx *gin.Context)
	PostEvent(ctx *gin.Context)
	GetHealth(ctx *gin.Context)
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
				apiV1.POST("/place", v1.PostPlace)
				apiV1.GET("/place/:id", v1.GetPlace)
				apiV1.GET("/place/:id/distance", v1.GetDistance)
			}
			{
				apiV1.POST("/event", v1.PostEvent)
			}
			{
				apiV1.GET("/health", v1.GetHealth)
			}
		}

	}

	// TODO swagger only for dev env
	srv.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return srv
}

func (s *service) Run() error {
	return s.engine.Run(fmt.Sprintf(":%d", s.cfg.Port))
}
