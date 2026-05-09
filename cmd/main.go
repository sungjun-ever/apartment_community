package main

import (
	"apart_community/config"
	"apart_community/database"
	"apart_community/registry"
	"apart_community/router"

	"github.com/gin-gonic/gin"
)

func main() {
	env := config.LoadEnv()

	db := database.ConnectToPostgres(env)
	rdb := database.ConnectToRedis(env)

	container := registry.NewContainer(db, rdb)

	r := gin.Default()

	r = router.SetUpRouter(r, container)

	r.Run()
}
