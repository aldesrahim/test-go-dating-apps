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

	secured := group.Group("/")
	secured.Use(middlewares.NewJWTMiddleware(loggedInUser).GetHandler())

	profile := secured.Group("/profile")
	{
		profileController := controllers.NewProfileController(loggedInUser)
		profile.GET("/", profileController.GetProfile)
		profile.PUT("/", profileController.UpdateProfile)
	}

	subs := secured.Group("/subscription")
	{
		subsController := controllers.NewSubsController(loggedInUser)
		subs.GET("/active", subsController.GetActive)
		subs.GET("/products", subsController.GetProducts)
		subs.POST("/buy", subsController.Buy)
	}

	swipes := secured.Group("/swipes")
	{
		swipeController := controllers.NewSwipeController(loggedInUser)
		swipes.GET("/", swipeController.GetSwipeable)
		swipes.POST("/", swipeController.Swipe)
	}
}
