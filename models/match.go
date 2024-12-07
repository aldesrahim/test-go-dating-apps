package models

import "database/sql"

type Match struct {
	Model
	FirstUserID  uint
	FirstUser    User `gorm:"foreignKey:FirstUserID"`
	SecondUserID uint
	SecondUser   User `gorm:"foreignKey:SecondUserID"`
	ClosedAt     sql.NullTime
	ClosedUserID uint
	ClosedUser   User `gorm:"foreignKey:ClosedUserID"`
}
