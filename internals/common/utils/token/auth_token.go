package token

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/user/domain"
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetAuthHeaderToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	return strings.TrimPrefix(authHeader, "Bearer ")
}

func CreateAccessToken(claims *domain.AccessClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func CreateRefreshToken(claims *domain.RefreshClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateAccessToken(tokenString string) (*domain.AccessClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errUtils.NewAppError(errors.New("잘못된 signing method"), 401, errUtils.A002, errUtils.LevelInfo)
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, errUtils.NewAppError(err, 401, errUtils.A002, errUtils.LevelInfo)
	}

	if claims, ok := token.Claims.(*domain.AccessClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errUtils.NewAppError(errors.New("유효하지 않은 토큰"), 401, errUtils.A002, errUtils.LevelInfo)
}
