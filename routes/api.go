package routes

import (
	"aldesrahim/test-go-dating-apps/controllers"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	group := router.Group("/api")

	auth := group.Group("/auth")
	auth.POST("/register", controllers.Register)
}
