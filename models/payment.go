package models

import "database/sql"

type Payment struct {
	Model
	UserID    uint
	User      User
	PaidAt    sql.NullTime
	Amount    float64
	Status    string
	Reference sql.NullString
}
