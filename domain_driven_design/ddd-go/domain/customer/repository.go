package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/irononet/ddd-go/aggregate"
)

var (
	// ErrCustomerNotFound is returned when a customer is not found 
	ErrCustomerNotFound = errors.New("the customer was not found in the repo") 

	ErrFailedToAddCustomer = errors.New("failed to add customer to the repo") 

	ErrUpdateCustomer = errors.New("faile to update the customer in the repo") 
)

type CustomerRepository interface{
	Get(uuid.UUID) (aggregate.Customer, error) 
	Add(aggregate.Customer) error 
	Update(aggregate.Customer) error 
}