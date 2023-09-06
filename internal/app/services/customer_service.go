package services

import (
	"golang-rest-api/internal/app/repositories"
	"golang-rest-api/internal/domain"
)

type CustomerService interface {
	Create(request domain.Customer) (*domain.Customer, error)
}

type customerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) CustomerService {
	return &customerService{customerRepository: customerRepository}
}

func (s *customerService) Create(requestBody domain.Customer) (*domain.Customer, error) {

	// validate := validator.New()

	// validate.VarWithValue(requestBody.Name, "", "contains")

	// if requestBody.Name == "" {
	// 	return nil, errors.New("nama harus diisi")
	// }

	// if requestBody.PhoneNumber == "" {
	// 	return nil, errors.New("phone harus diisi")
	// }

	// if requestBody.Address == "" {
	// 	return nil, errors.New("alamat harus diisi")
	// }

	customer, err := s.customerRepository.Create(&requestBody)

	if err != nil {
		panic(err)
	}

	return customer, nil
}
