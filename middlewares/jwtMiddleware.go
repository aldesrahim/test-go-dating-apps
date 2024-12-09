package middlewares

import (
	"aldesrahim/test-go-dating-apps/database"
	"aldesrahim/test-go-dating-apps/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JTWMiddleware struct {
	LoggedInUser *LoggedInUser
}

func NewJWTMiddleware(loggedInUser *LoggedInUser) *JTWMiddleware {
	return &JTWMiddleware{LoggedInUser: loggedInUser}
}

func (m *JTWMiddleware) GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Missing token",
			})
			c.Abort()

			return
		}

		tokenString := authHeader[len("Bearer "):]
		claims, err := utils.ValidateJWT(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Invalid token",
			})
			c.Abort()

			return
		}

		userID, ok := claims["id"].(float64)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Invalid token",
			})
			c.Abort()

			return
		}

		if err := database.DB.Where("id = ?", userID).First(&m.LoggedInUser.User).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  "Invalid user ID",
			})
			c.Abort()

			return
		}

		m.LoggedInUser.IsLoggedIn = true

		c.Next()
	}
}
