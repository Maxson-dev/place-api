package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func SetGlobalLogger(lvl string) {
	programLevel := new(slog.LevelVar)

	programLevel.Set(resolve(lvl))

	fmt.Println(programLevel.Level())

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel}))

	slog.SetDefault(logger)

}

func resolve(lvl string) slog.Level {
	switch lvl {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}
