package models

import (
	"time"
)

type Match struct {
	ID           uint       `gorm:"primarykey" json:"id"`
	FirstUserID  uint       `json:"first_user_id"`
	FirstUser    User       `gorm:"foreignKey:FirstUserID" json:"first_user"`
	SecondUserID uint       `json:"second_user_id"`
	SecondUser   User       `gorm:"foreignKey:SecondUserID" json:"second_user"`
	ClosedAt     *time.Time `json:"closed_at"`
	ClosedUserID uint       `json:"closed_user_id"`
	ClosedUser   User       `gorm:"foreignKey:ClosedUserID" json:"closed_user"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
