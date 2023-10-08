package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/irononet/ddd-go/aggregate"
	"github.com/irononet/ddd-go/domain/customer"
)

// MemoryRepository fulfills the CustomerRepository interface
type MemoryRepository struct{
	customers map[uuid.UUID]aggregate.Customer 
	sync.Mutex
}

// New is a factory function to generate a new repository of customers 
func New() *MemoryRepository{
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Get finds a customer by ID 
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error){
	if customer, ok := mr.customers[id]; ok{
		return customer, nil 
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

// Add will add a new customer to the repository 
func (mr *MemoryRepository) Add(c aggregate.Customer) error{
	if mr.customers == nil{
		// safety check if customers is not create, should'nt happen 
		// if using  
		mr.Lock() 
		mr.customers = make(map[uuid.UUID]aggregate.Customer) 
		mr.Unlock()
	}

	// make sure customer isn't already in the repository 
	if _, ok := mr.customers[c.GetId()]; ok{
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock() 
	mr.customers[c.GetId()] = c 
	mr.Unlock() 
	return nil 
}

// udpate will replace an existing customer information with the new 
// customer information 

func (mr *MemoryRepository) Update(c aggregate.Customer) error{
	// make sure Customer is in the repository 
	if _, ok := mr.customers[c.GetId()]; !ok{
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	mr.Lock() 
	mr.customers[c.GetId()] = c 
	mr.Unlock() 
	return nil 
}