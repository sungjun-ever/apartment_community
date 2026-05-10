package middleware

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/common/utils/token"
	"apart_community/internals/user/repository"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(redisAuthRepo repository.RedisAuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			_ = c.Error(errUtils.NewAppError(errors.New("인증 정보가 없습니다"), 401, "A003"))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			_ = c.Error(errUtils.NewAppError(errors.New("토큰 형식이 잘못됐습니다"), 401, "A002"))
			c.Abort()
			return
		}

		_, err := token.ValidateAccessToken(parts[1])

		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		storedToken, err := redisAuthRepo.GetSession(c, parts[1])

		if err != nil || storedToken == "" {
			_ = c.Error(errUtils.NewAppError(errors.New("인증이 만료됐습니다"), 401, "A001"))
			c.Abort()
			return
		}

		c.Next()
	}
}
