package database

import (
	"apart_community/config"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis(env *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.RedisHost + ":" + env.RedisPort,
		Password: "",
		DB:       0,
	})

	return rdb
}
