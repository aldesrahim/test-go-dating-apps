package models

type Interest struct {
	Model
	code  string `gorm:"uniqueIndex:code_group"`
	group string `gorm:"uniqueIndex:code_group"`
	label string
}
