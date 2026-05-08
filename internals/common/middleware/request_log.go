package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()
		
		traceID, _ := c.Get("trace_id")

		slog.Info("request",
			"ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"trace_id", traceID,
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"status", c.Writer.Status(),
			"latency", time.Since(start),
		)
	}
}
