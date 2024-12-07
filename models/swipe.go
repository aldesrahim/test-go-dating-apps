package models

import "time"

type Swipe struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user"`
	Counter   uint      `json:"counter"`
	Remaining uint      `json:"remaining"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
