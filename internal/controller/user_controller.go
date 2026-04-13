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
	users, err := u.us.GetAllUsers()

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

	getUser, err := u.us.GetUser(uri.ID)

	if err != nil {
		log.Println(err.Error())
		c.JSON(404, dto.ErrorResponse("사용자 정보가 없습니다."))
		return
	}

	userInfo := user.Resources{
		ID:             getUser.ID,
		UUID:           getUser.UUID,
		Email:          getUser.Email,
		CreatedAt:      getUser.CreatedAt,
		Nickname:       getUser.Profile.Nickname,
		Status:         getUser.Status,
		ProfileImageId: getUser.Profile.ProfileImageId,
	}

	c.JSON(200, dto.SuccessResponse(userInfo))
}

func (u *UserController) CreateUser(c *gin.Context) {
	var req user.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("user:", err.Error())
		c.JSON(400, dto.ErrorResponse("잘못된 요청입니다."))
		return
	}

	userEntity := req.ToUserEntity()
	profileEntity := req.ToProfileEntity()

	createUser, err := u.us.CreateUser(userEntity, profileEntity)

	if err != nil {
		log.Println(err.Error())
		c.JSON(500, dto.ErrorResponse("사용자 생성에 실패했습니다."))
		return
	}

	userInfo := user.Resources{
		ID:             createUser.ID,
		UUID:           createUser.UUID,
		Email:          createUser.Email,
		CreatedAt:      createUser.CreatedAt,
		Nickname:       createUser.Profile.Nickname,
		Status:         createUser.Status,
		ProfileImageId: createUser.Profile.ProfileImageId,
	}

	c.JSON(201, dto.SuccessResponse(userInfo))
}

func (u *UserController) UpdateUser(c *gin.Context) {
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

func (u *UserController) DeleteUser(c *gin.Context) {
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
