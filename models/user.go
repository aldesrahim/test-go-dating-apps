package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID            uint           `gorm:"primarykey" json:"ID"`
	Email         string         `gorm:"uniqueIndex" json:"email"`
	PasswordHash  string         `json:"password_hash"`
	Bio           sql.NullString `json:"bio"`
	DOB           sql.NullTime   `json:"dob"`
	Age           sql.NullInt16  `json:"age"`
	Name          sql.NullString `json:"name"`
	Gender        sql.NullString `json:"gender"`
	Interests     []Interest     `gorm:"many2many:user_interests;" json:"interests"`
	UserInterests []UserInterest `json:"user_interests"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (u *User) IsProfileCompleted() bool {
	return u.Name.Valid
}

type UserDTO struct {
	ID    uint
	Email string
	Name  string
}

func (u *User) ToDTO() UserDTO {
	return UserDTO{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name.String,
	}
}
