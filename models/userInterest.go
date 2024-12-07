package models

type UserInterest struct {
	UserID     uint   `gorm:"primaryKey" json:"user_id"`
	InterestID uint   `gorm:"primaryKey" json:"interest_id"`
	Value      string `json:"value"`
}
