package models

type UserInterest struct {
	UserID     uint `gorm:"primaryKey"`
	InterestID uint `gorm:"primaryKey"`
	Value      string
}
