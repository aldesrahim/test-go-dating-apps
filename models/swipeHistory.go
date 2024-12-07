package models

type SwipeHistory struct {
	Model
	SwipeID      uint
	UserID       uint
	User         User
	SwipedUserID uint
	SwipedUser   User `gorm:"foreignKey:SwipedUserID"`
}
