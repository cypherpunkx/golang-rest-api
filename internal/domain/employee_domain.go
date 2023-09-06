package domain

type Employee struct {
	ID   string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
}
