package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/irononet/ddd-go/aggregate"
	"github.com/irononet/ddd-go/domain/customer"
	"github.com/irononet/ddd-go/domain/product"
	"github.com/irononet/ddd-go/domain/customer/memory"

	prodmemory "github.com/irononet/ddd-go/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error 

type OrderService struct{
	customers customer.CustomerRepository
	products product.ProductRepository
}

// NewOrderService takes a variable amount of OrderConfiguation function 
// and return each Order 
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error){
	os := &OrderService{} 
	// Apply all configurations passed in 
	for _, cfg := range cfgs{
		err := cfg(os) 
		if err != nil{
			return nil, err 
		}
	}

	return os, nil 
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error){

	// Get the customer 
	c, err := o.customers.Get(customerID) 
	if err != nil{
		return 0, err 
	}

	var products []aggregate.Product 
	var price float64 

	for _, id := range productIDs{
		p, err := o.products.GetByID(id) 
		if err != nil{
			return 0, err 
		}

		products = append(products, p) 
		price += p.GetPrice()
	}

	// All Products exists in store, now we can create the order 
	log.Printf("Customer: %s has ordered %d products", c.GetId(), len(products))

	return price, nil 
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration{
	// return a function that matches the OrderConfiguration alias, 
	// you need to return this so that the parent function can 
	// take in all the  
	return func(os *OrderService) error{
		os.customers = cr 
		return nil 
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService 
func WithMemoryCustomerRepository() OrderConfiguration{
	// Create the memory repo, if we needed parameters, such as connection string 
	cr := memory.New() 
	return WithCustomerRepository(cr)
}


func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration{
	return func(os *OrderService) error{
		pr := prodmemory.New() 

		// Add items to repo 
		for _, p := range products{
			err := pr.Add(p) 
			if err != nil{
				return err 
			}
		}

		os.products = pr 
		return nil 
	}
}