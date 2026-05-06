package middleware

import (
	"apart_community/internals/errUtils"
	"apart_community/internals/response"
	"errors"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func ErrorLogMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			var appErr *errUtils.AppError
			traceID, _ := c.Get("trace_id")

			if errors.As(err, &appErr) {
				slog.Error("Application Error",
					"trace_id", traceID.(string),
					"code", appErr.Status,
					"message", appErr.Message,
					"error", appErr.Err,
				)
				response.Fail(c, appErr.Status, appErr.Code, appErr.Message)
			} else {
				slog.Error("Unknown Error",
					"trace_id", traceID.(string),
					"error", err,
				)
				response.Fail(c, 500, "S001", "INTERNAL_SERVER_ERROR")
			}
		}
	}
}
