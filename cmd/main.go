package main

import (
	"apart_community/config"
	"apart_community/internal/database"
	"apart_community/registry"
	"apart_community/router"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := database.ConnectMySQLDB()

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	container := registry.NewContainer(db)

	r := gin.Default()
	router.SetupRouter(r, container)
}
