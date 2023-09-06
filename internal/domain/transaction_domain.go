package domain

import "time"

type Transaction struct {
	ID       string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrderID  string    `gorm:"type:uuid;not null" json:"order_id"`
	BillDate time.Time `gorm:"type:date;not null" json:"bill_date"`
}
