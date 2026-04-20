package middleware

import (
	"apart_community/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-Trace-Id")

		if traceID == "" {
			traceID = uuid.New().String()
		}

		ctx := utils.WithTraceID(c, traceID)
		c.Request = c.Request.WithContext(ctx)

		c.Writer.Header().Set("X-Trace-Id", traceID)

		c.Next()
	}
}
