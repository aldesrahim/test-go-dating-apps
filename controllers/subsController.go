package controllers

import (
	"aldesrahim/test-go-dating-apps/database"
	"aldesrahim/test-go-dating-apps/middlewares"
	"aldesrahim/test-go-dating-apps/models"
	"aldesrahim/test-go-dating-apps/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type SubsController struct {
	LoggedInUser *middlewares.LoggedInUser
}

func NewSubsController(loggedInUser *middlewares.LoggedInUser) *SubsController {
	return &SubsController{LoggedInUser: loggedInUser}
}

func (ctrl *SubsController) GetActive(c *gin.Context) {
	userActiveSubs, err := services.GetUserActiveSubs(ctrl.LoggedInUser.User.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "User is not subscribed",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   userActiveSubs,
	})
}

func (ctrl *SubsController) GetProducts(c *gin.Context) {
	var products []models.Product

	database.DB.Preload("Periods", func(db *gorm.DB) *gorm.DB {
		return db.Order("product_periods.monthly_period ASC")
	}).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   products,
	})
}

func (ctrl *SubsController) Buy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
