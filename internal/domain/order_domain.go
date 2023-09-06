package domain

import (
	"time"
)

type Order struct {
	ID            string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CustomerID    string    `gorm:"type:uuid;not null" json:"customer_id"`
	OrderDetailID string    `gorm:"type:uuid;not null" json:"order_detail_id"`
	EntryDate     time.Time `gorm:"type:date;not null;default:CURRENT_DATE" json:"entry_date"`
	FinishDate    time.Time `gorm:"type:date;not null" json:"finish_date"`
}
