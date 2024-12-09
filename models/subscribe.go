package models

import "time"

type Subscribe struct {
	ID              uint           `gorm:"primarykey" json:"id"`
	UserID          uint           `json:"user_id"`
	User            *User          `json:"user,omitempty"`
	PaymentID       uint           `json:"payment_id"`
	Payment         *Payment       `json:"payment,omitempty"`
	ProductID       uint           `json:"product_id"`
	Product         *Product       `json:"product,omitempty"`
	ProductPeriodID uint           `json:"product_period_id"`
	ProductPeriod   *ProductPeriod `json:"product_period,omitempty"`
	StartDate       time.Time      `gorm:"type=DATE" json:"start_date"`
	EndDate         time.Time      `gorm:"type=DATE" json:"end_date"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}
