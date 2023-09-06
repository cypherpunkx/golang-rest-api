package domain

type OrderDetail struct {
	ID        string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProductID string `gorm:"type:uuid;not null" json:"product_id"`
	OrderID   string `gorm:"type:uuid;not null" json:"order_id"`
	Quantity  int    `gorm:"type:int;not null" json:"quantity"`
	Subtotal  int    `gorm:"type:int;not null;default:0" json:"subtotal"`
}
