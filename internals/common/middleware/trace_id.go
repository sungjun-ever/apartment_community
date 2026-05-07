package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get("X-Trace-ID")

		if traceID == "" {
			traceID = uuid.New().String()
		}

		ctx := context.WithValue(c, "trace_id", traceID)
		c.Request = c.Request.WithContext(ctx)

		c.Writer.Header().Set("X-Trace-ID", traceID)
		c.Next()
	}
}
