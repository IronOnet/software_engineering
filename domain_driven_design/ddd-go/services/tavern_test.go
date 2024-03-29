package services  

import (
	"testing" 

	"github.com/google/uuid" 
	"github.com/irononet/ddd-go/aggregate"
)


func Test_Tavern(t *testing.T){

	products := init_products(t) 

	os, err := NewOrderService(
		WithMemoryCustomerRepository(), 
		WithMemoryProductRepository(products), 
	)

	if err != nil{
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os)) 
	if err != nil{
		t.Error(err) 
	}

	cust, err := aggregate.NewCustomer("Percy") 
	if err != nil{
		t.Error(err)
	}

	err = os.customers.Add(cust) 
	if err != nil{
		t.Error(err) 
	}

	order := []uuid.UUID{
		products[0].GetID(), 
	}

	err = tavern.Order(cust.GetId(), order) 
	if err != nil{
		t.Error(err)
	}
}

func Test_MongoTavern(t *testing.T){
	products := init_products(t) 

	os, err := NewOrderService(
		WithMongoCustomerRepository("mongodb://localhost:27017"), 
		WithMemoryProductRepository(products), 
	)

	if err != nil{
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os)) 
	if err != nil{
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Percy") 
	if err != nil{
		t.Error(err)
	}

	err = os.customers.Add(cust) 
	if err != nil{
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(), 
	}

	err = tavern.Order(cust.GetId(), order) 
	if err != nil{
		t.Error(err) 
	}
}