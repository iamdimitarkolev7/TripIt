package routes

import (
	"tripit/controllers"
	"tripit/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	home := router.Group("/")
	{
		home.GET("/", controllers.GetHome())
	}

	authentication := router.Group("/user")
	{
		authentication.GET("/:id", middleware.Authentication(), controllers.GetUser())
		authentication.POST("/register", controllers.Register())
		authentication.POST("/login", controllers.Login())
		authentication.GET("/logout", middleware.Authentication(), controllers.Logout())
	}
}
