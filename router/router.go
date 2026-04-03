package router

import (
	"apart_community/registry"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, ct *registry.Container) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
