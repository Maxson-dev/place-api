package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Maxson-dev/place-api/config"
	"github.com/Maxson-dev/place-api/internal/controller"
	v1 "github.com/Maxson-dev/place-api/internal/controller/v1"
	filrepo "github.com/Maxson-dev/place-api/internal/infra/database/file"
	"github.com/Maxson-dev/place-api/internal/infra/s3"
	"github.com/Maxson-dev/place-api/internal/pkg/logger"
	fileusecase "github.com/Maxson-dev/place-api/internal/usecase/file"
	"github.com/Maxson-dev/place-api/migration"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	cfg := config.MustParse("./config.yml")

	logger.SetGlobalLogger(cfg.Logger.Level)

	slog.Debug("App init started")

	err := migration.Migrate(ctx, cfg.Postgres.Dsn)
	if err != nil {
		slog.Error("migration error: %s", err)
		os.Exit(1)
	}

	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	fileRepo := filrepo.New()

	s3Client, err := s3.New(s3.Config{
		Region:    cfg.S3.Region,
		Endpoint:  cfg.S3.Host,
		AccessKey: cfg.S3.AccessKey,
		SecretKey: cfg.S3.SecretKey,
	})
	if err != nil {
		slog.Error("s3 init error: %s", err)
		os.Exit(1)
	}

	fileUC := fileusecase.New(cfg.S3.Bucket, s3Client, fileRepo)

	v1api := v1.New(fileUC)

	httpcfg := controller.HTTPConfig{
		Port:                cfg.HTTP.Port,
		MaxMultipartSizeMib: cfg.HTTP.MaxMultipartSizeMib,
	}
	app := controller.New(engine, httpcfg, v1api)

	if err := app.Run(); err != nil {
		slog.Error("app run error: %s", err)
		os.Exit(1)
	}
}
