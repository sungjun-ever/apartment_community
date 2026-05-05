package router

import "github.com/gin-gonic/gin"

func SetUpRouter(router *gin.Engine) *gin.Engine {
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/login", func(c *gin.Context) {
				c.String(200, "this is login")
			})
			v1.POST("/register", func(c *gin.Context) {
				c.String(200, "this is register")
			})

			users := v1.Group("/users")
			{
				users.GET("/", func(c *gin.Context) {
					c.String(200, "this is get users")
				})

				users.GET("/:publicID", func(c *gin.Context) {
					c.String(200, "this is get user")
				})
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
