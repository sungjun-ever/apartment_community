package domain

import "github.com/golang-jwt/jwt/v5"

type AccessClaims struct {
	PublicID string `json:"publicId"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	//Roles    []string `json:"role"`

	jwt.RegisteredClaims
}

type RefreshClaims struct {
	PublicID string `json:"publicId"`
	jwt.RegisteredClaims
}

type UserAuthBase struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}
