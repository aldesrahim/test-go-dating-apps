package services

import (
	"aldesrahim/test-go-dating-apps/database"
	"aldesrahim/test-go-dating-apps/models"
	"errors"
	"time"
)

var ErrReachedTodaySwipeLimit = errors.New("reached today swipe limit")
var ErrSwipedAllTheUsers = errors.New("swiped all the users")

type SwipeAttribute struct {
	Swipeable *models.SwipeHistory
	Status    uint8
}

func GetTodaySwipe(user *models.User) *models.Swipe {
	now := time.Now()

	var swipe *models.Swipe

	if err := database.DB.Preload("User").Where("user_id = ? AND DATE(created_at) = ?", user.ID, now.Format(time.DateOnly)).First(&swipe).Error; err != nil {
		var remaining uint
		userActiveSubs := GetUserActiveSubs(user.ID)

		if userActiveSubs == nil || !userActiveSubs.IsUnlimitedSwipes() {
			remaining = 10
		}

		swipe = &models.Swipe{
			UserID:    user.ID,
			User:      user,
			Counter:   0,
			Remaining: &remaining,
		}

		database.DB.Save(&swipe)
	}

	return swipe
}

func GetTodaySwipeable(swipe *models.Swipe) (swipeable *models.SwipeHistory, err error) {
	if swipe.Remaining != nil && *swipe.Remaining == 0 {
		return nil, ErrReachedTodaySwipeLimit // You've reached today's limit
	}

	if err := database.DB.Preload("SwipedUser").Where("swipe_id = ? AND status IS NULL", swipe.ID).First(&swipeable).Error; err != nil {
		var newUser models.User

		subQuery := database.DB.Select("1").Where("swipe_histories.swipe_id = ? AND swipe_histories.swiped_user_id = users.id", swipe.ID).Table("swipe_histories")
		err = database.DB.Where("id <> ? AND gender = ? AND interest = ? AND NOT EXISTS (?)", swipe.UserID, swipe.User.Interest, swipe.User.Gender, subQuery).Order("RAND()").First(&newUser).Error

		if err != nil {
			return nil, ErrSwipedAllTheUsers // You're already swiped all the registered users
		}

		swipeable = &models.SwipeHistory{
			SwipeID:      swipe.ID,
			Swipe:        swipe,
			UserID:       swipe.UserID,
			SwipedUserID: newUser.ID,
			SwipedUser:   &newUser,
		}

		database.DB.Save(&swipeable)
	}

	return swipeable, nil
}

func CounterTodaySwipe(swipe *models.Swipe, attribute *SwipeAttribute) (err error) {
	if swipe.Remaining != nil && *swipe.Remaining == 0 {
		return ErrReachedTodaySwipeLimit // You've reached today's limit
	}

	swipe.Counter++

	if swipe.Remaining != nil {
		newRemaining := *swipe.Remaining - uint(1)
		swipe.Remaining = &newRemaining
	}

	if err = database.DB.Save(&swipe).Error; err != nil {
		return err
	}

	attribute.Swipeable.Status = &attribute.Status

	database.DB.Save(&attribute.Swipeable)

	return nil
}

func IsSwipeableVerified(swipeable *models.SwipeHistory) bool {
	userActiveSubs := GetUserActiveSubs(swipeable.SwipedUserID)

	if userActiveSubs == nil {
		return false
	}

	return userActiveSubs.IsUserVerified()
}
