package controller

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/common/response"
	"apart_community/internals/common/utils/token"
	"apart_community/internals/user/domain"
	"apart_community/internals/user/service"
	"time"

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
		_ = c.Error(errUtils.NewAppError(err, 400, errUtils.C001, errUtils.LevelInfo))
		return
	}

	user, err := ac.as.Auth(c, uab)

	if err != nil {
		_ = c.Error(err)
		return
	}

	refreshDuration := time.Hour * 24 * 7

	accessToken, refreshToken, err := ac.as.IssueToken(user, refreshDuration)

	if err != nil {
		_ = c.Error(err)
		return
	}

	session := domain.UserSession{
		RefreshToken: *refreshToken,
		IP:           c.ClientIP(),
		UserAgent:    c.Request.UserAgent(),
		CreatedAt:    time.Now().Format(time.DateTime),
	}

	err = ac.as.CreateSession(c, user.PublicID, &session, refreshDuration)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.SetCookie("refreshToken", *refreshToken, int(refreshDuration.Seconds()), "/", "", false, true)

	response.OK(c, 200, gin.H{
		"access_token": accessToken,
	}, nil)
}

func (ac *AuthController) Logout(c *gin.Context) {
	tokenString := token.GetAuthHeaderToken(c)
	claims, err := token.ValidateAccessToken(tokenString)

	if err != nil {
		_ = c.Error(err)
		return
	}

	err = ac.as.DestroySession(c, tokenString, claims)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.SetCookie("refreshToken", "", -1, "/", "", false, true)
	response.OK(c, 200, nil, nil)
}
