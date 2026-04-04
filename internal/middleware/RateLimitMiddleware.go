package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimitMiddleware 메모리 기반 -> 이후에 레디스로?
func RateLimitMiddleware() gin.HandlerFunc {
	rate := limiter.Rate{
		Limit:  10,
		Period: 1 * time.Second,
	}

	store := memory.NewStore()
	instance := limiter.New(store, rate)

	return mgin.NewMiddleware(instance)
}
