package utils

import (
	"context"
	"log/slog"
)

func InfoLogWithContext(ctx context.Context, message, traceId string) {
	slog.InfoContext(ctx, message, "trace_id:", traceId)
}

func ErrorLogWithContext(ctx context.Context, message, methodName, traceId string) {
	slog.ErrorContext(ctx, message, "method:", methodName, "trace_id:", traceId)
}
