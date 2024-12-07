package models

import "database/sql"

type SubscribeDetail struct {
	Model
	UserID          uint
	User            User
	ProductID       uint
	Product         Product
	ProductPeriodID uint
	ProductPeriod   ProductPeriod
	StartDate       sql.NullTime
	EndDate         sql.NullTime
}
