package services

import (
	"aldesrahim/test-go-dating-apps/database"
	"aldesrahim/test-go-dating-apps/models"
	"time"
)

func GetUserActiveSubs(userID uint) (models.Subscribe, error) {
	today := time.Now().Format(time.DateOnly)
	var activeSubs models.Subscribe

	err := database.DB.Preload("User").Preload("Product").Where("user_id = ?", userID).Where("start_date <= ? AND end_date >= ?", today, today).First(&activeSubs).Error

	return activeSubs, err
}
