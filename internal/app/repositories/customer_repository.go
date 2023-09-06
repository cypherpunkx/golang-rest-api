package repositories

import (
	"golang-rest-api/internal/domain"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *domain.Customer) (*domain.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(customer *domain.Customer) (*domain.Customer, error) {
	if err := r.db.Create(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}
