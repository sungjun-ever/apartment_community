package main

import (
	"apart_community/config"
	"apart_community/database"
	"apart_community/registry"
	"apart_community/router"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := database.ConnectToPostgres()
	rdb := database.ConnectToRedis()

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(2 * time.Minute)

	container := registry.NewContainer(db, rdb)

	r := gin.Default()

	r = router.SetUpRouter(r, container)

	r.Run()
}
