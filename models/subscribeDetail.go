package models

import (
	"database/sql"
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
	StartDate       sql.NullTime  `json:"start_date"`
	EndDate         sql.NullTime  `json:"end_date"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}
