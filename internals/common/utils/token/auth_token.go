package token

import (
	"apart_community/internals/user/domain"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(claims *domain.AccessClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func CreateRefreshToken(claims *domain.RefreshClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
