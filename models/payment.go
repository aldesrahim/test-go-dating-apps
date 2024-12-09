package models

import (
	"time"
)

type Payment struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	UserID    uint       `json:"user_id"`
	User      User       `json:"user"`
	PaidAt    *time.Time `json:"paid_at"`
	Amount    float64    `json:"amount"`
	Status    string     `json:"status"`
	Reference *string    `json:"reference"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
