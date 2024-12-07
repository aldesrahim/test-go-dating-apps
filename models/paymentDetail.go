package models

type PaymentDetail struct {
	Model
	PaymentID       uint
	Payment         Payment
	ProductID       uint
	Product         Product
	ProductPeriodID uint
	ProductPeriod   ProductPeriod
	MonthlyPeriod   uint
	Price           float64
}
