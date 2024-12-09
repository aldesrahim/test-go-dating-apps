package models

import (
	"time"
)

type Payment struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	UserID          uint           `json:"user_id"`
	User            *User          `json:"user,omitempty"`
	ProductID       uint           `json:"product_id"`
	Product         *Product       `json:"product,omitempty"`
	ProductPeriodID uint           `json:"product_period_id"`
	ProductPeriod   *ProductPeriod `json:"product_period,omitempty"`
	MonthlyPeriod   uint           `json:"monthly_period"`
	Price           float64        `json:"price"`
	PaidAt          *time.Time     `json:"paid_at"`
	Status          string         `json:"status"`
	Reference       *string        `json:"reference"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}
