package router

import (
	"apart_community/internals/common/middleware"
	"apart_community/registry"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine, ct *registry.Container) *gin.Engine {
	router.Use(middleware.RateLimitMiddleware(ct.Redis))
	router.Use(middleware.TraceIdMiddleware())
	router.Use(middleware.RequestLogMiddleware())
	router.Use(middleware.ErrorLogMiddleWare())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/login", func(c *gin.Context) {
				c.String(200, "this is login")
			})

			v1.POST("/register", ct.UserController.StoreUser)

			users := v1.Group("/users")
			{
				users.GET("/", ct.UserController.GetUsers)

				users.GET("/:publicID", ct.UserController.GetUser)

				users.PUT("/:publicID", func(c *gin.Context) {
					c.String(200, "this is update user")
				})
				users.DELETE("/:publicID", func(c *gin.Context) {
					c.String(200, "this is delete user")
				})
			}

			aparts := v1.Group("/apartments")
			{
				aparts.GET("/", func(c *gin.Context) {
					c.String(200, "this is get aparts")
				})
				aparts.GET("/:publicID", func(c *gin.Context) {
					c.String(200, "this is get apart")
				})
				aparts.PUT("/:publicID", func(c *gin.Context) {
					c.String(200, "this is update apart")
				})
				aparts.DELETE("/:publicID", func(c *gin.Context) {
					c.String(200, "this is delete apart")
				})
			}
		}
	}

	return router
}
