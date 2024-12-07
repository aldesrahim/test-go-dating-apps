package database

import "aldesrahim/test-go-dating-apps/models"

func Migrate() {
	DB.AutoMigrate(
		&models.Interest{},
		&models.Product{},
		&models.ProductPeriod{},
		&models.User{},
		&models.UserInterest{},
		&models.Swipe{},
		&models.SwipeHistory{},
		&models.Match{},
		&models.Payment{},
		&models.PaymentDetail{},
		&models.Subscribe{},
		&models.SubscribeDetail{},
	)
}
