package main

import (
	"apart_community/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r = router.SetUpRouter(r)

	r.Run()
}
