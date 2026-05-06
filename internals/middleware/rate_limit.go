package middleware

import (
	"apart_community/internals/errUtils"
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
			_ = c.Error(errUtils.NewAppError(err, 500, "S001"))
			c.Abort()
			return
		}

		if count == 1 {
			rdb.Expire(ctx, key, time.Minute)
		}

		if count > 60 {
			_ = c.Error(errUtils.NewAppError(err, 429, "C004"))
			c.Abort()
			return
		}

		c.Next()
	}
}
