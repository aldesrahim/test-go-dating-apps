package models

import "time"

type Product struct {
	ID        uint             `gorm:"primarykey" json:"id"`
	Name      string           `json:"name"`
	Code      string           `gorm:"uniqueIndex" json:"code"`
	Periods   *[]ProductPeriod `gorm:"ProductID" json:"periods,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
