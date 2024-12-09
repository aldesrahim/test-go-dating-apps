package controllers

import (
	"aldesrahim/test-go-dating-apps/database"
	"aldesrahim/test-go-dating-apps/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ProfileController struct {
	LoggedInUser *middlewares.LoggedInUser
}

func NewProfileController(loggedInUser *middlewares.LoggedInUser) *ProfileController {
	return &ProfileController{LoggedInUser: loggedInUser}
}

func (ctrl *ProfileController) GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user":   ctrl.LoggedInUser.User,
	})
}
func (ctrl *ProfileController) UpdateProfile(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Bio      string `json:"bio" binding:"required"`
		DOB      string `json:"dob" binding:"required,datetime=2006-01-02"`
		Location string `json:"location" binding:"required"`
		Gender   string `json:"gender" binding:"required,oneof=male female other"`
		Interest string `json:"interest" binding:"required,oneof=male female other"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})

		return
	}

	dob, _ := time.Parse(time.DateOnly, input.DOB)

	user := ctrl.LoggedInUser.User
	user.Name = &input.Name
	user.Bio = &input.Bio
	user.DOB = &dob
	user.Location = &input.Location
	user.Gender = &input.Gender
	user.Interest = &input.Interest

	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user":   user,
	})
}
