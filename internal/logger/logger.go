package logger

import (
	"log/slog"
	"os"

	"github.com/pkg/errors"
)

func New(env string) (*slog.Logger, error) {
	const op = "logger.New"
	var log *slog.Logger
	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		return nil, errors.Wrap(errors.New("application deployment environment is not defined"), op)
	}

	return log, nil
}
