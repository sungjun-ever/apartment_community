package router

import (
	"apart_community/internal/middleware"
	"apart_community/registry"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, ct *registry.Container) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.RateLimitMiddleware())
	{
		v1.POST("/register", ct.UserController.StoreUser)

		users := v1.Group("/users")
		{
			users.GET("/", ct.UserController.GetUsers)
			users.GET("/:id", ct.UserController.GetUser)
			users.PUT("/:id", ct.UserController.EditUser)
			users.DELETE(":id", ct.UserController.DestroyUser)

			users.PUT("/:id/apartment", ct.UserController.EditBelongApart)
		}

	}
}
