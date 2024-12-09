package controllers

import (
	"aldesrahim/test-go-dating-apps/database"
	"aldesrahim/test-go-dating-apps/middlewares"
	"aldesrahim/test-go-dating-apps/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	LoggedInUser *middlewares.LoggedInUser
}

func NewAuthController(LoggedInUser *middlewares.LoggedInUser) *AuthController {
	return &AuthController{LoggedInUser: LoggedInUser}
}

func invalidEmailOrPassword(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"error":  "The email or password you entered is incorrect",
	})
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	ctrl.LoggedInUser.Reset()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})

		return
	}

	hashedPassword, _ := utils.HashPassword(input.Password)

	user := ctrl.LoggedInUser.User

	user.Email = input.Email
	user.PasswordHash = hashedPassword

	if err := database.DB.Create(&user).Error; err != nil {
		if utils.ErrDBIsDuplicate(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "The email you entered is already registered",
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

func (ctrl *AuthController) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	ctrl.LoggedInUser.Reset()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})

		return
	}

	user := ctrl.LoggedInUser.User

	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if utils.ErrDBIsNotFound(err) {
			invalidEmailOrPassword(c)

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})

		return
	}

	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		invalidEmailOrPassword(c)

		return
	}

	token, _ := utils.GenerateJWT(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user":   user,
		"token":  token,
	})
}
