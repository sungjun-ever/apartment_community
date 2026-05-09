package main

import (
	"apart_community/config"
	"apart_community/database"
	"apart_community/registry"
	"apart_community/router"

	"github.com/gin-gonic/gin"
)

func main() {
	flags := config.ParseFlags()
	env := config.LoadEnv(flags)

	db := database.ConnectToPostgres(env)
	rdb := database.ConnectToRedis(env)

	database.Seeder(db, flags)

	container := registry.NewContainer(db, rdb)

	r := gin.Default()

	r = router.SetUpRouter(r, container)

	r.Run()
}
