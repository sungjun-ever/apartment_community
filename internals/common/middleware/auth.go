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
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			_ = c.Error(errUtils.NewAppError(errors.New("인증 정보가 없습니다"), 401, errUtils.A003, errUtils.LevelInfo))
			c.Abort()
			return
		}

		accessTokenString := token.GetAuthHeaderToken(c)

		isBlackListToken, err := redisAuthRepo.IsBlacklisted(c, accessTokenString)

		if isBlackListToken == 1 {
			_ = c.Error(errUtils.NewAppError(errors.New("인증이 만료됐습니다"), 401, errUtils.A001, errUtils.LevelInfo))
			c.Abort()
			return
		}

		if accessTokenString == "" {
			_ = c.Error(errUtils.NewAppError(errors.New("토큰 형식이 잘못됐습니다"), 401, errUtils.A002, errUtils.LevelInfo))
			c.Abort()
			return
		}

		claims, err := token.ValidateAccessToken(accessTokenString)

		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		storedToken, err := redisAuthRepo.GetSession(c, claims.PublicID)

		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		if storedToken == "" {
			_ = c.Error(errUtils.NewAppError(errors.New("인증이 만료됐습니다"), 401, errUtils.A001, errUtils.LevelInfo))
			c.Abort()
			return
		}

		c.Next()
	}
}
