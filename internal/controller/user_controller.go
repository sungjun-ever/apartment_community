package controller

import (
	"apart_community/internal/dto"
	"apart_community/internal/dto/user"
	"apart_community/internal/model"
	"apart_community/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us service.UserService
}

func NewUserController(us service.UserService) *UserController {
	return &UserController{us: us}
}

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.us.FindAllUsers()

	if err != nil {
		log.Println(err.Error())
		c.JSON(404, dto.ErrorResponse("사용자 정보가 없습니다."))
		return
	}

	c.JSON(200, dto.SuccessResponse(users))
}

func (u *UserController) GetUser(c *gin.Context) {
	var uri user.UriRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		log.Println(err.Error())
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	getUser, err := u.us.FindUser(uri.ID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(404, dto.ErrorResponse("사용자 정보가 없습니다."))
		return
	}

	userInfo := user.NewResource(getUser)

	c.JSON(200, dto.SuccessResponse(userInfo))
}

func (u *UserController) StoreUser(c *gin.Context) {
	var req user.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("user:", err.Error())
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	userEntity := req.ToUserEntity()
	profileEntity := req.ToProfileEntity()

	if userEntity == nil || profileEntity == nil {
		log.Println("StoreUser entity is empty", userEntity, profileEntity)
		c.JSON(500, dto.ErrorResponse("사용자 생성에 실패했습니다."))
		return
	}

	createUser, err := u.us.CreateUser(userEntity, profileEntity)

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, dto.ErrorResponse("사용자 생성에 실패했습니다."))
		return
	}

	userInfo := user.NewResource(createUser)

	c.JSON(201, dto.SuccessResponse(userInfo))
}

func (u *UserController) EditUser(c *gin.Context) {
	var uUser model.User

	if err := c.ShouldBindJSON(&uUser); err != nil {
		log.Println(err.Error())
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	updateUser, err := u.us.UpdateUser(uUser)

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, dto.ErrorResponse("사용자 수정에 실패했습니다."))
		return
	}

	c.JSON(200, dto.SuccessResponse(updateUser))
}

func (u *UserController) EditBelongApart(c *gin.Context) {
	var ubaRequest model.UserBelongApartment

	if err := c.ShouldBindBodyWithJSON(&ubaRequest); err != nil {
		log.Println(err.Error())
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	err := u.us.UpdateBelongApart(ubaRequest)

	if err != nil {
		log.Println(err.Error())
	}

	c.JSON(201, dto.SuccessEmptyResponse[model.UserBelongApartment]())
}

func (u *UserController) DestroyUser(c *gin.Context) {
	var uri user.UriRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		log.Println(err.Error())
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	err := u.us.DeleteUser(uri.ID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, dto.ErrorResponse("사용자 삭제에 실패했습니다."))
		return
	}

	c.JSON(200, dto.SuccessResponse("사용자 삭제에 성공했습니다."))
}
