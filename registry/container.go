package registry

import (
	"apart_community/internal/controller"
	"apart_community/internal/repository"
	"apart_community/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	UserController *controller.UserController
}

func NewContainer(db *gorm.DB) *Container {
	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	userSvc := service.NewUserService(userRepo, profileRepo, db)

	return &Container{
		UserController: controller.NewUserController(*userSvc),
	}
}
