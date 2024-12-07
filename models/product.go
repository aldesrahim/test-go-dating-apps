package models

type Product struct {
	Model
	Name string
	Code string `gorm:"uniqueIndex"`
}
