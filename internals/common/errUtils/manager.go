package errUtils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
)

type ErrorManager struct {
	logger *slog.Logger
}

func NewErrorManager(filepath string) *ErrorManager {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	handler := slog.NewJSONHandler(f, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	return &ErrorManager{
		logger: slog.New(handler),
	}
}

func (m *ErrorManager) Handle(ctx context.Context, err error) {
	if err == nil {
		return
	}

	var appErr *AppError
	ok := errors.As(err, &appErr)

	traceID := ctx.Value("trace_id")

	if !ok {
		m.logger.Error("UNKNOWN_ERROR", "trace_id", traceID, "error", err)
		return
	}

	logAttrs := []any{
		slog.String("trace_id", traceID.(string)),
		slog.String("code", appErr.Code.String()),
		slog.String("message", appErr.Code.GetMessage()),
	}

	switch appErr.Level {
	case LevelFatal:
		m.logger.Error("Application Error", logAttrs...)
		sendAlert(appErr)
	case LevelWarn:
		m.logger.Warn("Application Error", logAttrs...)
	default:
		m.logger.Info("Application Error", logAttrs...)
	}
}

func sendAlert(err *AppError) {
	fmt.Println("send something")
}
