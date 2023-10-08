package aggregate 

import (
	"errors" 
	
	"github.com/google/uuid"
	"github.com/irononet/ddd-go/entity" 
	"github.com/irononet/ddd-go/valueobject" 
) 



var (
	ErrInvalidPerson = errors.New("a customer has to be a valid person") 
	
)

type Customer struct{
	// person is the root entity of a customer 
	// which means the person.ID is the main idenfier for 
	// this aggregate 
	person *entity.Person 
	// a customer can hold many products 
	products []*entity.Item 
	// a customer can perform many transactions 
	transactions []*valueobject.Transaction
}


func NewCustomer(name string) (Customer, error){
	// check that the name is non empty 
	if name == ""{
		return Customer{}, ErrInvalidPerson 
	}

	person := &entity.Person{
		Name : name, 
		ID: uuid.New(), 
	}

	return Customer{
		person: person, 
		products: make([]*entity.Item, 0), 
		transactions: make([]*valueobject.Transaction, 0),
	}, nil 
}


func (c *Customer) GetId() uuid.UUID{
	return c.person.ID
}

// setID sets the root ID 
func (c *Customer) SetID(id uuid.UUID){
	if c.person == nil{
		c.person = &entity.Person{}
	}

	c.person.ID = id 
}

// SetName changes the name of the Customer 
func (c *Customer) SetName(name string){
	if c.person == nil{
		c.person = &entity.Person{}
	}

	c.person.Name = name 
}


// SetName changes the name of the Customer 
func (c *Customer) GetName() string{
	return c.person.Name
}
