package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	ID    uint
	Email string
	Role  string
	jwt.RegisteredClaims
}

func getSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func Sign(userClaims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	secretKey := getSecretKey()

	if len(secretKey) == 0 {
		return "", errors.New("암호화 키가 없습니다")
	}

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	parsedToken, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if _, ok := parsedToken.Claims.(*UserClaims); ok && parsedToken.Valid {
		return tokenString, nil
	}

	return "", err
}
