package routes

import (
	"aldesrahim/test-go-dating-apps/controllers"
	"aldesrahim/test-go-dating-apps/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	loggedInUser := middlewares.NewLoggedInUser()

	group := router.Group("/api")

	auth := group.Group("/auth")
	{
		authController := controllers.NewAuthController(loggedInUser)
		auth.POST("/register", authController.Register)
		auth.POST("/", authController.Login)
	}

	profile := group.Group("/profile")
	profile.Use(middlewares.NewJWTMiddleware(loggedInUser).GetHandler())
	{
		profile.GET("/", controllers.NewProfileController(loggedInUser).GetProfile)
		profile.PUT("/", controllers.NewProfileController(loggedInUser).UpdateProfile)
	}
}
