package product 

import (
	"errors" 

	"github.com/google/uuid" 
	"github.com/irononet/ddd-go/aggregate"
)

var (
	// ErrorProductNotFound is returned when a product is not found 
	ErrProductNotFound = errors.New("the product was not found") 

	ErrProductAlreadyExist = errors.New("the product already exists") 
)

// ProductRepository is the repository interface to fulfill to use the product 
type ProductRepository interface{
	GetAll() ([]aggregate.Product, error) 
	GetByID(id uuid.UUID) (aggregate.Product, error) 
	Add(product aggregate.Product) error 
	Update(product aggregate.Product) error 
	Delete(id uuid.UUID) error 
}