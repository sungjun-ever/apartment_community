package registry

import (
	"apart_community/internal/controller"
	"apart_community/internal/repository"
	"apart_community/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	UserController *controller.UserController
	AuthController *controller.AuthController
}

func NewContainer(db *gorm.DB) *Container {
	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)
	belongApartRepo := repository.NewBelongApartRepository(db)

	userSvc := service.NewUserService(userRepo, profileRepo, belongApartRepo, db)

	return &Container{
		UserController: controller.NewUserController(*userSvc),
		AuthController: controller.NewAuthController(*userSvc),
	}
}
