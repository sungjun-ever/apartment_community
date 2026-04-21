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
		c.JSON(404, dto.ErrorResponse("사용자 정보가 없습니다."))
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
		utils.ErrorLogWithContext(ctx, err.Error(), "GetUser", traceID)
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	getUser, err := u.us.FindUser(ctx, uri.ID)

	if err != nil {
		c.JSON(404, dto.ErrorResponse("사용자 정보가 없습니다."))
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
		utils.ErrorLogWithContext(ctx, err.Error(), "StoreUser", traceID)
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	createUser, err := u.us.CreateUser(ctx, req)

	if err != nil {
		c.JSON(500, dto.ErrorResponse("사용자 생성에 실패했습니다."))
		return
	}

	userInfo := user.NewResource(createUser)

	c.JSON(201, dto.SuccessResponse(userInfo))
}

func (u *UserController) EditUser(c *gin.Context) {
	ctx := c.Request.Context()
	traceID := utils.GetTraceID(ctx)

	utils.InfoLogWithContext(ctx, "start EditUser", traceID)

	var profile user.ProfileRequest

	if err := c.ShouldBindJSON(&profile); err != nil {
		utils.ErrorLogWithContext(ctx, err.Error(), "EditUser", traceID)
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	updateUser, err := u.us.UpdateUser(ctx, profile)

	if err != nil {
		c.JSON(500, dto.ErrorResponse("사용자 수정에 실패했습니다."))
		return
	}

	userInfo := user.NewResource(updateUser)

	c.JSON(200, dto.SuccessResponse(userInfo))
}

func (u *UserController) EditBelongApart(c *gin.Context) {
	ctx := c.Request.Context()
	traceID := utils.GetTraceID(ctx)

	utils.InfoLogWithContext(ctx, "start EditBelongApart", traceID)

	var ubaRequest model.UserBelongApartment

	if err := c.ShouldBindBodyWithJSON(&ubaRequest); err != nil {
		utils.ErrorLogWithContext(ctx, err.Error(), "EditBelongApart", traceID)
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	err := u.us.UpdateBelongApart(ctx, ubaRequest)

	if err != nil {
		c.JSON(500, dto.ErrorResponse("사용자 아파트 등록에 실패했습니다."))
	}

	c.JSON(201, dto.SuccessEmptyResponse[model.UserBelongApartment]())
}

func (u *UserController) DestroyUser(c *gin.Context) {
	ctx := c.Request.Context()
	traceID := utils.GetTraceID(ctx)

	utils.InfoLogWithContext(ctx, "start DestroyUser", traceID)

	var uri user.UriRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		utils.ErrorLogWithContext(ctx, err.Error(), "DestroyUser", traceID)
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	err := u.us.DeleteUser(ctx, uri.ID)

	if err != nil {
		c.JSON(500, dto.ErrorResponse("사용자 삭제에 실패했습니다."))
		return
	}

	c.JSON(200, dto.SuccessEmptyResponse[*model.User]())
}
