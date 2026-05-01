package utils

import (
	"context"
	"log/slog"
)

func InfoLogWithContext(ctx context.Context, message string) {
	slog.InfoContext(ctx, message, "trace_id:", GetTraceID(ctx))
}

func ErrorLogWithContext(ctx context.Context, message, methodName string) {
	slog.ErrorContext(ctx, message, "method:", methodName, "trace_id:", GetTraceID(ctx))
}
