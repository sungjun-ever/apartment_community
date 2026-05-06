package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func RateLimitMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		key := "rate_limit:" + c.ClientIP()

		count, err := rdb.Incr(ctx, key).Result()

		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
			return
		}

		if count == 1 {
			rdb.Expire(ctx, key, time.Minute)
		}

		if count > 60 {
			c.AbortWithStatusJSON(429, gin.H{"message": "Too Many Request"})
			return
		}

		c.Next()
	}
}
