package registry

import (
	"apart_community/internals/user/controller"
	"apart_community/internals/user/repository"
	"apart_community/internals/user/service"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Container struct {
	Postgres       *gorm.DB
	Redis          *redis.Client
	AuthRepo       repository.AuthRepository
	UserController *controller.UserController
	AuthController *controller.AuthController
}

func NewContainer(db *gorm.DB, rdb *redis.Client) *Container {
	authRepo := repository.NewAuthRepository(rdb)
	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)

	userSvc := service.NewService(userRepo, profileRepo, db)
	authSvc := service.NewAuthService(authRepo, userRepo, db, rdb)

	return &Container{
		Postgres:       db,
		Redis:          rdb,
		AuthRepo:       authRepo,
		UserController: controller.NewUserController(*userSvc),
		AuthController: controller.NewAuthController(*authSvc),
	}
}
