package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Maxson-dev/place-api/config"
	"github.com/Maxson-dev/place-api/internal/controller"
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

	fileUC := fileusecase.New(fileusecase.Config{})

	app := controller.New(engine, controller.HTTPConfig{
		Port:                cfg.HTTP.Port,
		MaxMultipartSizeMib: cfg.HTTP.MaxMultipartSizeMib,
	})

	if err := app.Run(); err != nil {
		slog.Error("app run error: %s", err)
		os.Exit(1)
	}
}
