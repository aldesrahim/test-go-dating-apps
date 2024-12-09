package services

import (
	"aldesrahim/test-go-dating-apps/database"
	"aldesrahim/test-go-dating-apps/models"
	"database/sql"
	"time"
)

func GetUserActiveSubs(userID uint) *models.Subscribe {
	today := time.Now().Format(time.DateOnly)
	var activeSubs *models.Subscribe

	if err := database.DB.Preload("Product").Where("user_id = ?", userID).Where("start_date <= @today AND end_date >= @today", sql.Named("today", today)).First(&activeSubs).Error; err != nil {
		return nil
	}

	return activeSubs
}

func UserSubsPaymentSuccess(referenceNumber string) models.Subscribe {
	var payment models.Payment

	if err := database.DB.Preload("ProductPeriod").Where("reference = ?", referenceNumber).First(&payment).Error; err != nil {
		panic(err)
	}

	monthlyPeriod := payment.ProductPeriod.MonthlyPeriod

	now := time.Now()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endDate := startDate.AddDate(0, int(monthlyPeriod), 1)

	subscribe := models.Subscribe{
		UserID:          payment.UserID,
		PaymentID:       payment.ID,
		ProductID:       payment.ProductID,
		ProductPeriodID: payment.ProductPeriodID,
		StartDate:       startDate,
		EndDate:         endDate,
	}

	database.DB.Save(&subscribe)

	return subscribe
}
