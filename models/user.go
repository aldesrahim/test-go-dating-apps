package models

import (
	"database/sql"
)

type User struct {
	Model
	Email         string `gorm:"uniqueIndex"`
	PasswordHash  string
	Bio           sql.NullString
	DOB           sql.NullTime
	Age           sql.NullInt16
	Name          sql.NullString
	Gender        sql.NullString
	Interests     []Interest `gorm:"many2many:user_interests;"`
	UserInterests []UserInterest
}
