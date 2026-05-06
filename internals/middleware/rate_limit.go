package middleware

import (
	"apart_community/internals/response"
	"apart_community/internals/utils"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RateLimitMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		key := utils.RateLimitKey(c.ClientIP())

		count, err := rdb.Incr(ctx, key).Result()

		if err != nil {
			response.AbortWithError(c, 500, "S001", "INTERNAL_SERVER_ERROR")
			return
		}

		if count == 1 {
			rdb.Expire(ctx, key, time.Minute)
		}

		if count > 60 {
			response.AbortWithError(c, 429, "C004", "TOO_MANY_REQUEST")
			return
		}

		c.Next()
	}
}
