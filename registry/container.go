package registry

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Container struct {
}

func NewContainer(db *gorm.DB, rdb *redis.Client) *Container {

	return &Container{}
}
