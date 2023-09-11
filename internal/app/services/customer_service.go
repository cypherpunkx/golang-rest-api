package services

import (
	"golang-rest-api/internal/app/repositories"
	"golang-rest-api/internal/domain"
)

type CustomerService interface {
	Create(request *domain.Customer) (*domain.Customer, error)
	View(request *[]domain.Customer) (*[]domain.Customer, error)
}

type customerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) CustomerService {
	return &customerService{customerRepository: customerRepository}
}

func (s *customerService) Create(requestBody *domain.Customer) (*domain.Customer, error) {
	return s.customerRepository.Create(requestBody)
}

func (s *customerService) View(requestBody *[]domain.Customer) (*[]domain.Customer, error) {
	return s.customerRepository.View(requestBody)
}
