package models

import "time"

type PaymentDetail struct {
	ID              uint          `gorm:"primarykey" json:"id"`
	PaymentID       uint          `json:"payment_id"`
	Payment         Payment       `json:"payment"`
	ProductID       uint          `json:"product_id"`
	Product         Product       `json:"product"`
	ProductPeriodID uint          `json:"product_period_id"`
	ProductPeriod   ProductPeriod `json:"product_period"`
	MonthlyPeriod   uint          `json:"monthly_period"`
	Price           float64       `json:"price"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
