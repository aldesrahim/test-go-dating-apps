package models

import "time"

type SwipeHistory struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	SwipeID      uint      `json:"swipe_id"`
	Swipe        *Swipe    `json:"swipe,omitempty"`
	UserID       uint      `json:"user_id"`
	User         *User     `json:"user,omitempty"`
	SwipedUserID uint      `json:"swiped_user_id"`
	SwipedUser   *User     `gorm:"foreignKey:SwipedUserID" json:"swiped_user,omitempty"`
	Status       *uint8    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
