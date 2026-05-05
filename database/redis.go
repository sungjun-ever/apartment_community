package database

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func ConnectToRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	return rdb
}
