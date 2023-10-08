package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/irononet/ddd-go/domain/customer/mongo"
)

type TavernConfiguraton func(os *Tavern) error 

type Tavern struct{
	OrderService *OrderService 

	BillingService interface{}
}

// NewTavern takes a variable amount of TavernConfigurations and builds 
// a tavern 
func NewTavern(cfgs ...TavernConfiguraton) (*Tavern, error){
	// Create the tavern 
	t := &Tavern{} 

	// Apply all configurations passed in 
	for _, cfg := range cfgs{
		// Pass the service into the configuration function 
		err := cfg(t) 
		if err != nil{
			return nil, err 
		}
	}

	return t, nil 
}

// WithOrderService applies a given OrderService to the Tavern 
func WithOrderService(os *OrderService) TavernConfiguraton{
	return func(t *Tavern) error{
		t.OrderService = os 
		return nil 
	}
}

// Order performs an order for a customer 
func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error{
	price, err := t.OrderService.CreateOrder(customer, products) 
	if err != nil{
		return err 
	}
	log.Printf("Bill the Customer: %0.f", price) 

	// Bill the customer 
	return nil 
}

func WithMongoCustomerRepository(connectionString string) OrderConfiguration{
	return func(os *OrderService) error{
		cr, err := mongo.New(context.Background(), connectionString)
		if err != nil{
			return err 
		}
		os.customers = cr 
		return nil 
	}
}