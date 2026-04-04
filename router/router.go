package router

import (
	"apart_community/registry"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, ct *registry.Container) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", ct.UserController.CreateUser)

		users := v1.Group("/users")
		{
			users.GET("/", ct.UserController.GetUsers)
			users.GET("/:id", ct.UserController.GetUser)
			users.PUT("/:id", ct.UserController.UpdateUser)
			users.DELETE(":id", ct.UserController.DeleteUser)
		}

	}
}
