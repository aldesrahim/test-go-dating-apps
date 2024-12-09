package models

import (
	"time"
)

type Payment struct {
	ID              uint          `gorm:"primarykey" json:"id"`
	UserID          uint          `json:"user_id"`
	User            User          `json:"user"`
	ProductID       uint          `json:"product_id"`
	Product         Product       `json:"product"`
	ProductPeriodID uint          `json:"product_period_id"`
	ProductPeriod   ProductPeriod `json:"product_period"`
	MonthlyPeriod   uint          `json:"monthly_period"`
	Price           float64       `json:"price"`
	PaidAt          *time.Time    `json:"paid_at"`
	Amount          float64       `json:"amount"`
	Status          string        `json:"status"`
	Reference       *string       `json:"reference"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
