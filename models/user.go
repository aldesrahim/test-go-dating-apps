package models

import (
	"aldesrahim/test-go-dating-apps/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint       `gorm:"primarykey" json:"id"`
	Email        string     `gorm:"uniqueIndex" json:"email"`
	PasswordHash string     `gorm:"type:text" json:"-"`
	Bio          *string    `gorm:"type:text" json:"bio"`
	DOB          *time.Time `json:"dob"`
	Location     *string    `gorm:"type:text" json:"location"`
	Age          *uint16    `json:"age"`
	Name         *string    `json:"name"`
	Gender       *string    `json:"gender"`
	Interest     *string    `json:"interest"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (u *User) IsProfileCompleted() bool {
	return u.Name != nil
}

func (u *User) BeforeSave(*gorm.DB) (err error) {
	if u.DOB != nil {
		age := uint16(utils.GetAgeFromDate(*u.DOB))
		u.Age = &age
	}

	return
}
