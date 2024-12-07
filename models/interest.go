package models

import "time"

type Interest struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Code      string    `gorm:"uniqueIndex:code_group" json:"code"`
	Group     string    `gorm:"uniqueIndex:code_group" json:"group"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
