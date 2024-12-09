package controllers

import (
	"aldesrahim/test-go-dating-apps/database"
	"aldesrahim/test-go-dating-apps/models"
	"aldesrahim/test-go-dating-apps/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})

		return
	}

	hashedPassword, _ := utils.HashPassword(input.Email)
	user := models.User{
		Email:        input.Email,
		PasswordHash: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if utils.ErrorDBIsDuplicate(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "Email already exists",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})

		return
	}

	token, _ := utils.GenerateJWT(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user":   user,
		"token":  token,
	})
}
