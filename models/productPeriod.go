package models

import "time"

type ProductPeriod struct {
	ID            uint      `gorm:"primarykey" json:"id"`
	ProductID     uint      `json:"product_id"`
	Product       Product   `json:"product"`
	MonthlyPeriod uint      `json:"monthly_period"`
	Price         float64   `json:"price"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
