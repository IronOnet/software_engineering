package services 

import (
	"testing" 

	"github.com/google/uuid" 
	"github.com/irononet/ddd-go/aggregate"
)

func init_products(t *testing.T) []aggregate.Product{
	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 1.99) 
	if err != nil{
		t.Error(err) 
	}

	peanuts, err := aggregate.NewProduct("Peanuts", "Healthy Snacks", 0.99) 
	if err != nil{
		t.Error(err) 
	}

	wine, err := aggregate.NewProduct("Wine", "Healthy snacks", 0.99) 
	if err != nil{
		t.Error(err) 
	}

	products := []aggregate.Product{
		beer, peanuts, wine, 
	}

	return products 
}

func TestOrder_NewOrderService(t *testing.T){
	// Create a few products to insert into a memory repo 
	products := init_products(t) 

	os, err := NewOrderService(
		WithMemoryCustomerRepository(), 
		WithMemoryProductRepository(products), 
	)

	if err != nil{
		t.Error(err) 
	}

	// Add Customer 
	cust, err := aggregate.NewCustomer("Percy") 
	if err != nil{
		t.Error(err)
	}

	err = os.customers.Add(cust) 
	if err != nil{
		t.Error(err) 
	}

	// Perform Order for one beer 
	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetId(), order) 

	if err != nil{
		t.Error(err)
	}
}