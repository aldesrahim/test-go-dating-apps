package models

import (
	"time"
)

type Match struct {
	ID           uint       `gorm:"primarykey" json:"id"`
	FirstUserID  *uint      `json:"first_user_id"`
	FirstUser    *User      `gorm:"foreignKey:FirstUserID" json:"first_user,omitempty"`
	SecondUserID *uint      `json:"second_user_id"`
	SecondUser   *User      `gorm:"foreignKey:SecondUserID" json:"second_user,omitempty"`
	ClosedAt     *time.Time `json:"closed_at"`
	ClosedUserID *uint      `json:"closed_user_id"`
	ClosedUser   *User      `gorm:"foreignKey:ClosedUserID" json:"closed_user,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (m *Match) IsMatched() bool {
	return m.FirstUserID != nil && m.SecondUserID != nil
}
