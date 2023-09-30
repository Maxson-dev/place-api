package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Maxson-dev/place-api/config"
	"github.com/Maxson-dev/place-api/internal/controller"
	v1 "github.com/Maxson-dev/place-api/internal/controller/v1"
	filrepo "github.com/Maxson-dev/place-api/internal/infra/database/file"
	db "github.com/Maxson-dev/place-api/internal/infra/database/pgx-wrapper"
	placerepo "github.com/Maxson-dev/place-api/internal/infra/database/place"
	"github.com/Maxson-dev/place-api/internal/infra/s3"
	"github.com/Maxson-dev/place-api/internal/pkg/logger"
	fileuc "github.com/Maxson-dev/place-api/internal/usecase/file"
	placeuc "github.com/Maxson-dev/place-api/internal/usecase/place"
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

	// ____________INFRA______________

	masterNode, err := db.New(
		ctx,
		cfg.Postgres.Dsn,
		db.Config{
			PoolMax: cfg.Postgres.PoolMax,
		},
	)
	if err != nil {
		slog.Error("db init error: %s", err)
		os.Exit(1)
	}

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

	// ____________REPO______________

	fileRepo := filrepo.New()
	placeRepo := placerepo.New()

	// ____________USECASE______________

	fileUC := fileuc.New(
		masterNode,
		s3Client,
		fileRepo,
		fileuc.Config{
			StorageBucket:          cfg.S3.Bucket,
			DownloadURLLifetimeMin: cfg.S3.DownloadURLLifetimeMin,
		},
	)

	placeUC := placeuc.New(masterNode, placeRepo)

	// ____________CONTROLLER______________

	v1api := v1.New(fileUC, placeUC)

	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	app := controller.New(
		engine,
		v1api,
		controller.HTTPConfig{
			Port:                cfg.HTTP.Port,
			MaxMultipartSizeMib: cfg.HTTP.MaxMultipartSizeMib,
		},
	)

	if err := app.Run(); err != nil {
		slog.Error("app run error: %s", err)
		os.Exit(1)
	}
}
