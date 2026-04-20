package controller

import (
	"apart_community/internal/dto"
	"apart_community/internal/dto/user"
	"apart_community/internal/model"
	"apart_community/internal/service"
	"apart_community/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us service.UserService
}

func NewUserController(us service.UserService) *UserController {
	return &UserController{us: us}
}

func (u *UserController) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	traceID := utils.GetTraceID(ctx)
	utils.InfoLogWithContext(ctx, "start GetUsers", traceID)

	users, err := u.us.FindAllUsers(ctx)

	if err != nil {
		msg := "사용자 정보가 없습니다."
		utils.ErrorLogWithContext(ctx, "GetUsers", msg, traceID)
		c.JSON(404, dto.ErrorResponse(msg))
		return
	}

	c.JSON(200, dto.SuccessResponse(users))
}

func (u *UserController) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	traceID := utils.GetTraceID(ctx)
	utils.InfoLogWithContext(ctx, "start GetUser", traceID)

	var uri user.UriRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		msg := "잘못된 요청입니다."
		utils.ErrorLogWithContext(ctx, msg, "GetUser", traceID)
		c.JSON(400, dto.ErrorResponse(msg))
		return
	}

	getUser, err := u.us.FindUser(ctx, uri.ID)

	if err != nil {
		msg := "사용자 정보가 없습니다."
		utils.ErrorLogWithContext(ctx, msg, "GetUser", traceID)
		c.JSON(404, dto.ErrorResponse(msg))
		return
	}

	userInfo := user.NewResource(getUser)

	c.JSON(200, dto.SuccessResponse(userInfo))
}

func (u *UserController) StoreUser(c *gin.Context) {
	ctx := c.Request.Context()
	traceID := utils.GetTraceID(ctx)

	utils.InfoLogWithContext(ctx, "start StoreUser", traceID)

	var req user.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		msg := "잘못된 요청입니다."
		utils.ErrorLogWithContext(ctx, msg, "StoreUser", traceID)
		c.JSON(400, dto.ErrorResponse(msg))
		return
	}

	createUser, err := u.us.CreateUser(ctx, req)

	if err != nil {
		msg := "사용자 생성에 실패했습니다."
		utils.ErrorLogWithContext(ctx, msg, "StoreUser", traceID)
		c.JSON(500, dto.ErrorResponse(msg))
		return
	}

	userInfo := user.NewResource(createUser)

	c.JSON(201, dto.SuccessResponse(userInfo))
}

func (u *UserController) EditUser(c *gin.Context) {
	ctx := c.Request.Context()
	traceID := utils.GetTraceID(ctx)

	utils.InfoLogWithContext(ctx, "start EditUser", traceID)

	var uUser model.User

	if err := c.ShouldBindJSON(&uUser); err != nil {
		msg := "잘못된 요청입니다."
		utils.ErrorLogWithContext(ctx, msg, "EditUser", traceID)
		c.JSON(400, dto.ErrorResponse(msg))
		return
	}

	updateUser, err := u.us.UpdateUser(ctx, uUser)

	if err != nil {
		msg := "사용자 수정에 실패했습니다."
		utils.ErrorLogWithContext(ctx, msg, "EditUser", traceID)
		c.JSON(500, dto.ErrorResponse(msg))
		return
	}

	c.JSON(200, dto.SuccessResponse(updateUser))
}

func (u *UserController) EditBelongApart(c *gin.Context) {
	ctx := c.Request.Context()
	traceID := utils.GetTraceID(ctx)

	utils.InfoLogWithContext(ctx, "start EditBelongApart", traceID)

	var ubaRequest model.UserBelongApartment

	if err := c.ShouldBindBodyWithJSON(&ubaRequest); err != nil {
		msg := "잘못된 요청입니다."
		utils.ErrorLogWithContext(ctx, msg, "EditBelongApart", traceID)
		c.JSON(400, dto.ErrorResponse(msg))
		return
	}

	err := u.us.UpdateBelongApart(ctx, ubaRequest)

	if err != nil {
		msg := "사용자 아파트 등록에 실패했습니다."
		utils.ErrorLogWithContext(ctx, msg, "EditBelongApart", traceID)
		c.JSON(500, dto.ErrorResponse(msg))
	}

	c.JSON(201, dto.SuccessEmptyResponse[model.UserBelongApartment]())
}

func (u *UserController) DestroyUser(c *gin.Context) {
	ctx := c.Request.Context()
	traceID := utils.GetTraceID(ctx)

	utils.InfoLogWithContext(ctx, "start DestroyUser", traceID)

	var uri user.UriRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		msg := "잘못된 요청입니다."
		utils.ErrorLogWithContext(ctx, msg, "DestroyUser", traceID)
		c.JSON(400, dto.ErrorResponse(msg))
		return
	}

	err := u.us.DeleteUser(ctx, uri.ID)

	if err != nil {
		msg := "사용자 삭제에 실패했습니다."
		utils.ErrorLogWithContext(ctx, msg, "DestroyUser", traceID)
		c.JSON(500, dto.ErrorResponse(msg))
		return
	}

	c.JSON(200, dto.SuccessEmptyResponse[*model.User]())
}
