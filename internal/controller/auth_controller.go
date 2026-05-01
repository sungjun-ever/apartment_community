package controller

import (
	"apart_community/internal/dto"
	"apart_community/internal/dto/user"
	"apart_community/internal/service"
	"apart_community/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	us service.UserService
}

func NewAuthController(
	us service.UserService,
) *AuthController {
	return &AuthController{
		us: us,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	ctx := c.Request.Context()
	utils.InfoLogWithContext(ctx, "start Login")

	var lq user.LoginRequest

	if err := c.ShouldBindJSON(&lq); err != nil {
		utils.ErrorLogWithContext(ctx, err.Error(), "Login")
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	user, err := ac.us.GetUserByEmail(ctx, lq.Email)

	if err != nil {
		c.JSON(404, dto.ErrorResponse("사용자 정보가 없습니다"))
		return
	}

	userClaims := utils.UserClaims{
		user.ID,
		user.Email,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token, err := utils.Sign(userClaims)

	if err != nil {
		utils.ErrorLogWithContext(ctx, err.Error(), "Login")
		c.JSON(500, dto.ErrorResponse("토큰 생성에 실패했습니다."))
		return
	}

	c.JSON(200, dto.SuccessResponse(token))
}
