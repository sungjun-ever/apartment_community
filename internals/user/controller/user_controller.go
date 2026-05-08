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

func (uc *UserController) GetUsers(c *gin.Context) {
	var pr domain.PaginationRequest

	if err := c.ShouldBindQuery(&pr); err != nil {
		_ = c.Error(errUtils.NewAppError(err, 400, "C001"))
		return
	}

	users, total, err := uc.us.FindUsers(c, pr)

	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OK(c, 200, domain.NewUserResources(users), &response.Meta{
		Page:  pr.Page,
		Size:  pr.Size,
		Total: total,
	})
}

func (uc *UserController) GetUser(c *gin.Context) {
	var pidRequest domain.PublicIdUriRequest

	if err := c.ShouldBindUri(&pidRequest); err != nil {
		_ = c.Error(errUtils.NewAppError(err, 400, "C001"))
		return
	}

	user, err := uc.us.FindUser(c, pidRequest)

	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OK(c, 200, domain.NewUserResource(user), nil)
}

func (uc *UserController) StoreUser(c *gin.Context) {
	var rq domain.RegisterRequest

	if err := c.ShouldBindJSON(&rq); err != nil {
		_ = c.Error(errUtils.NewAppError(err, 400, "C001"))
		return
	}

	user, err := uc.us.CreateUser(c, rq)

	if err != nil {
		_ = c.Error(err)
		return
	}

	response.OK(c, 201, domain.NewUserResource(user), nil)
}
