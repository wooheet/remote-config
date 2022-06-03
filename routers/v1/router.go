package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wooheet/remote-config/handler"
)

func ApplyRoutes(r *gin.RouterGroup) {
	test := r.Group("/ping")
	{
		test.GET("/", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}
	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/logout", handler.Logout)
		auth.POST("/signup", handler.Signup)
	}
	config := r.Group("/config")
	{
		config.GET("/token", handler.Retrieve)
		config.POST("/token", handler.Registry)
		config.PUT("/token", handler.Update)
		config.PATCH("/token", handler.Update)
		config.DELETE("/token", handler.Delete)
		config.GET("/script-tag", handler.ScriptTag)
	}
}
