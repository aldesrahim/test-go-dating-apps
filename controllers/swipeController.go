package controllers

import (
	"aldesrahim/test-go-dating-apps/middlewares"
	"aldesrahim/test-go-dating-apps/models"
	"aldesrahim/test-go-dating-apps/services"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SwipeController struct {
	LoggedInUser *middlewares.LoggedInUser
}

func NewSwipeController(loggedInUser *middlewares.LoggedInUser) *SwipeController {
	return &SwipeController{LoggedInUser: loggedInUser}
}

func parseGeneralError(c *gin.Context, err error) {
	if err != nil && errors.Is(err, services.ErrReachedTodaySwipeLimit) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "You've reached today's limit",
		})
		c.Abort()

		return
	}

	if err != nil && errors.Is(err, services.ErrSwipedAllTheUsers) {
		c.JSON(466, gin.H{
			"status": 466,
			"error":  "Congratulations! You're already swiped all the registered users",
		})
		c.Abort()

		return
	}
}

func (ctrl *SwipeController) GetSwipeable(c *gin.Context) {
	swipe := services.GetTodaySwipe(ctrl.LoggedInUser.User)
	swipeable, err := services.GetTodaySwipeable(swipe)

	if err != nil {
		parseGeneralError(c, err)

		if !c.IsAborted() {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
		}

		return
	}

	var data = struct {
		Swipeable  *models.SwipeHistory `json:"swipeable"`
		IsVerified bool                 `json:"is_verified"`
	}{
		Swipeable:  swipeable,
		IsVerified: services.IsSwipeableVerified(swipeable),
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}

func (ctrl *SwipeController) Swipe(c *gin.Context) {
	swipe := services.GetTodaySwipe(ctrl.LoggedInUser.User)
	swipeable, err := services.GetTodaySwipeable(swipe)

	if err != nil {
		parseGeneralError(c, err)

		if !c.IsAborted() {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
		}

		return
	}

	var input struct {
		Status int `json:"status" binding:"required,oneof=0 1"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})

		return
	}

	attribute := services.SwipeAttribute{
		Swipeable: swipeable,
		Status:    uint8(input.Status),
	}

	err = services.CounterTodaySwipe(swipe, &attribute)

	if err != nil {
		parseGeneralError(c, err)

		if !c.IsAborted() {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
		}

		return
	}

	var data = struct {
		Swipeable  *models.SwipeHistory `json:"swipeable"`
		IsVerified bool                 `json:"is_verified"`
	}{
		Swipeable:  swipeable,
		IsVerified: services.IsSwipeableVerified(swipeable),
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
	})
}
