package models

import (
	"time"
)

type User struct {
	ID            uint           `gorm:"primarykey" json:"id"`
	Email         string         `gorm:"uniqueIndex" json:"email"`
	PasswordHash  string         `gorm:"type:text" json:"-"`
	Bio           *string        `gorm:"type:text" json:"bio"`
	DOB           *time.Time     `json:"dob"`
	Location      *string        `gorm:"type:text" json:"location"`
	Age           *uint16        `json:"age"`
	Name          *string        `json:"name"`
	Gender        *string        `json:"gender"`
	Interests     []Interest     `gorm:"many2many:user_interests;" json:"interests"`
	UserInterests []UserInterest `json:"user_interests"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (u *User) IsProfileCompleted() bool {
	return u.Name != nil
}
