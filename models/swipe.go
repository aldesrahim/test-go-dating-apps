package models

type Swipe struct {
	Model
	UserID    uint
	User      User
	Counter   uint
	Remaining uint
}
