package models

import (
	"time"
)

type SubscribeDetail struct {
	ID              uint          `gorm:"primarykey" json:"id"`
	UserID          uint          `json:"user_id"`
	User            User          `json:"user"`
	ProductID       uint          `json:"product_id"`
	Product         Product       `json:"product"`
	ProductPeriodID uint          `json:"product_period_id"`
	ProductPeriod   ProductPeriod `json:"product_period"`
	StartDate       *time.Time    `json:"start_date"`
	EndDate         *time.Time    `json:"end_date"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
