package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func getParsedToken(tokenString string) (*jwt.Token, error) {
	secretKey := getSecretKey()

	if len(secretKey) == 0 {
		return nil, errors.New("암호화 키가 없습니다")
	}

	parsedToken, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return parsedToken, nil
}

func CheckAdminRole(tokenString string) (bool, error) {
	parsedToken, err := getParsedToken(tokenString)

	if err != nil {
		return false, err
	}

	if claims, ok := parsedToken.Claims.(*UserClaims); ok && parsedToken.Valid {
		if claims.Role == "admin" {
			return true, nil
		}

		return false, errors.New("권한이 없습니다")
	}

	return false, err
}
