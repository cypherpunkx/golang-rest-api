package domain

type Customer struct {
	ID          string `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	Name        string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	PhoneNumber string `gorm:"type:varchar(100);not null" json:"phoneNumber" validate:"required"`
	Address     string `gorm:"type:varchar(100);not null" json:"address" validate:"required"`
}
