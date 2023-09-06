package domain

type Product struct {
	ID    string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name  string `gorm:"type:varchar(100);unique;not null" json:"name"`
	Price int    `gorm:"type:int;not null;default:0" json:"price"`
	Unit  string `gorm:"type:varchar(10);not null" json:"unit"`
}
