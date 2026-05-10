package controller

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/common/response"
	"apart_community/internals/user/domain"
	"apart_community/internals/user/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	as service.AuthService
}

func NewAuthController(as service.AuthService) *AuthController {
	return &AuthController{
		as: as,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	var uab domain.UserAuthBase

	if err := c.ShouldBindJSON(&uab); err != nil {
		_ = c.Error(errUtils.NewAppError(err, 400, "C001"))
		return
	}

	user, err := ac.as.Auth(c, uab)

	if err != nil {
		_ = c.Error(err)
		return
	}

	accessToken, refreshToken, err := ac.as.IssueToken(user)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.SetCookie("refreshToken", *refreshToken, 3600*24*7, "/", "", false, true)

	response.OK(c, 200, gin.H{
		"access_token": accessToken,
	}, nil)
}
