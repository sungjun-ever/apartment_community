package controller

import (
	"apart_community/internals/common/errUtils"
	"apart_community/internals/common/response"
	"apart_community/internals/user/domain"
	"apart_community/internals/user/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us *service.UserService
}

func NewUserController(us service.UserService) *UserController {
	return &UserController{us: &us}
}

func (uc *UserController) StoreUser(c *gin.Context) {
	var rq domain.RegisterRequest

	if err := c.ShouldBindJSON(&rq); err != nil {
		_ = c.Error(errUtils.NewAppError(err, 400, "C001"))
		return
	}

	user, err := uc.us.CreateUser(c, rq)

	if err != nil {
		_ = c.Error(errUtils.NewAppError(err, 500, "S001"))
		return
	}

	response.OK(c, 201, user, nil)
}
