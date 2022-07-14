package customer_usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type CustomerRegistrationUseCase interface {
	CreateCustomer(customer []*model.Customer) error
}

type customerRegistrationUseCase struct {
	customerRepo repository.CustomerRepository
}

func (c *customerRegistrationUseCase) CreateCustomer(customer []*model.Customer) error {
	return c.customerRepo.Create(customer)
}

func NewCustomerRegistrationUseCase(customerRepo repository.CustomerRepository) CustomerRegistrationUseCase {
	usecase := new(customerRegistrationUseCase)
	usecase.customerRepo = customerRepo
	return usecase
}
