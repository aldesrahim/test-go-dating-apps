package models

type ProductPeriod struct {
	Model
	ProductID     uint
	Product       Product
	MonthlyPeriod uint
	Price         float64
}
