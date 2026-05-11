package middleware

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/common/utils/token"
	"apart_community/internals/user/repository"
	"errors"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(redisAuthRepo repository.RedisAuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				_ = c.Error(errUtils.NewAppError(err, 500, "S001"))
				c.Abort()
			}
		}()

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			_ = c.Error(errUtils.NewAppError(errors.New("인증 정보가 없습니다"), 401, "A003"))
			c.Abort()
			return
		}

		accessTokenString := token.GetAuthHeaderToken(c)

		isBlackListToken, err := redisAuthRepo.IsBlacklisted(c, accessTokenString)

		if isBlackListToken == 1 {
			_ = c.Error(errUtils.NewAppError(errors.New("인증이 만료됐습니다"), 401, "A001"))
			c.Abort()
			return
		}

		if accessTokenString == "" {
			_ = c.Error(errUtils.NewAppError(errors.New("토큰 형식이 잘못됐습니다"), 401, "A002"))
			c.Abort()
			return
		}

		claims, err := token.ValidateAccessToken(accessTokenString)

		storedToken, err := redisAuthRepo.GetSession(c, claims.PublicID)

		if storedToken == "" {
			_ = c.Error(errUtils.NewAppError(errors.New("인증이 만료됐습니다"), 401, "A001"))
			c.Abort()
			return
		}

		c.Next()
	}
}
