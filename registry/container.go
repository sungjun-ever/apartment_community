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
	RedisAuthRepo  repository.RedisAuthRepository
	UserController *controller.UserController
	AuthController *controller.AuthController
}

func NewContainer(db *gorm.DB, rdb *redis.Client) *Container {
	redisAuthRepo := repository.NewRedisAuthRepository(rdb)
	userRepo := repository.NewGormUserRepository(db)
	profileRepo := repository.NewGormProfileRepository(db)

	userSvc := service.NewService(userRepo, profileRepo, db)
	authSvc := service.NewAuthService(redisAuthRepo, userRepo, db, rdb)

	return &Container{
		Postgres:       db,
		Redis:          rdb,
		RedisAuthRepo:  redisAuthRepo,
		UserController: controller.NewUserController(*userSvc),
		AuthController: controller.NewAuthController(*authSvc),
	}
}
