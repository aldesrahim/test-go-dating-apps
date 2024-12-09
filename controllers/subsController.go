package controllers

import (
	"aldesrahim/test-go-dating-apps/database"
	"aldesrahim/test-go-dating-apps/enums"
	"aldesrahim/test-go-dating-apps/middlewares"
	"aldesrahim/test-go-dating-apps/models"
	"aldesrahim/test-go-dating-apps/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type SubsController struct {
	LoggedInUser *middlewares.LoggedInUser
}

func NewSubsController(loggedInUser *middlewares.LoggedInUser) *SubsController {
	return &SubsController{LoggedInUser: loggedInUser}
}

func (ctrl *SubsController) GetActive(c *gin.Context) {
	userActiveSubs := services.GetUserActiveSubs(ctrl.LoggedInUser.User.ID)

	if userActiveSubs == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "You are currently not subscribe to anything",
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
	var input struct {
		ProductPeriodID uint `json:"product_period_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})

		return
	}

	productComboCode := string(enums.ProductTypeCombo1)
	user := ctrl.LoggedInUser.User
	userActiveSubs := services.GetUserActiveSubs(user.ID)

	// User not allowed to combine "Product Combo" with another product
	if userActiveSubs != nil && userActiveSubs.Product.Code == productComboCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "You can't combine your currently active subscription \"Combo\" with another product",
		})

		return
	}

	var validProduct *models.ProductPeriod

	if err := database.DB.Preload("Product").First(&validProduct, input.ProductPeriodID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "Invalid product period ID",
		})

		return
	}

	// User not allowed to combine "Product Combo" with another product
	if userActiveSubs != nil && validProduct.Product.Code == productComboCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "You can't combine your currently active subscription with \"Combo\" product",
		})

		return
	}

	// User not allowed to subscribe to the same subscription
	if userActiveSubs != nil && userActiveSubs.ProductID == validProduct.ProductID {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  fmt.Sprintf("You are already subscribed to %s until %s", userActiveSubs.Product.Name, userActiveSubs.EndDate.Format(time.DateOnly)),
		})

		return
	}

	referenceNumber := "TRX-001." + strconv.FormatInt(time.Now().Unix(), 10)

	payment := models.Payment{
		UserID:          user.ID,
		ProductID:       validProduct.ProductID,
		Product:         validProduct.Product,
		ProductPeriodID: validProduct.ID,
		MonthlyPeriod:   validProduct.MonthlyPeriod,
		Price:           validProduct.Price,
		Reference:       &referenceNumber, // this reference number should be from payment provider/gateways
		Status:          enums.PaymentStatusPending,
	}

	database.DB.Save(&payment)

	/* This should not exist here, because we need a proper
	payment validation after user paid (callbacks) */
	go func(referenceNumber string) {
		time.Sleep(10 * time.Second) // simulate user payment

		services.UserSubsPaymentSuccess(referenceNumber)
	}(referenceNumber)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   payment,
	})
}
