package registry

import (
	"apart_community/internals/user/controller"
	"apart_community/internals/user/domain"
	"apart_community/internals/user/service"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Container struct {
	Postgres       *gorm.DB
	Redis          *redis.Client
	UserController *controller.UserController
}

func NewContainer(db *gorm.DB, rdb *redis.Client) *Container {
	userRepo := domain.NewGormUserRepository(db)
	profileRepo := domain.NewGormProfileRepository(db)

	userSvc := service.NewService(userRepo, profileRepo, db)

	return &Container{
		Postgres:       db,
		Redis:          rdb,
		UserController: controller.NewUserController(*userSvc),
	}
}
