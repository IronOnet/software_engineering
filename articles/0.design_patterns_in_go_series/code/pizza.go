package main

import (
	"errors"
	"fmt"
)

type Pizza interface {
	Prepare() interface{}
	Bake() interface{}
	Cut() interface{}
	Box() interface{}
}

type PizzaStore interface {
	CreatePizza(pizzaType string) (Pizza, error)
}

const (
	KINSHASA_PIZZA_STORE = iota
	GOMA_PIZZA_STORE
	NEW_YORK_PIZZA_STORE
	JOHANESBURG_PIZZA_STORE
	ROME_PIZZA_STORE
)

const (
	VEGGIE_PIZZA_FACTORY = iota
	CLAM_PIZZA_FACTORY
	CHEESE_PIZZA_FACTORY
	PEPPERONI_PIZZA_FACTORY
)

type KinshasaStore struct{}

const (
	VEGGIE_PIZZA    = "VEGGIE_PIZZA"
	CLAM_PIZZA      = "CLAM_PIZZA"
	PEPPERONI_PIZZA = "PEPPERONI_PIZZA"
	CHEESE_PIZZA    = "CHEESE_PIZZA"
)

type VeggiePizza struct {
	Name string 
}

func (v *VeggiePizza) Prepare() interface{} {
	fmt.Println("Preparing veggie pizza")
	v.Name = "Veggie Pizza"
	return v 
}

func (v *VeggiePizza) Bake() interface{}{
	fmt.Println("Baking veggie pizza")
	return v
}

func (v *VeggiePizza) Cut() interface{}{
	fmt.Println("Cutting veggie pizza")
	return v
}

func (v *VeggiePizza) Box() interface{}{
	fmt.Println("Boxing veggie pizza") 
	fmt.Println("Pizza ready to be shipped..")
	return v
}

type ClamPizza struct{
	Name string 
}

func (c *ClamPizza) Prepare() interface{}{
	fmt.Println("Preparing Clam Pizza") 
	c.Name = "Clam Pizza"
	return c
}

func (c *ClamPizza) Bake() interface{}{
	fmt.Println("Baking Clam Pizza") 
	return c 
}

func (c *ClamPizza) Cut() interface{}{
	fmt.Println("Cutting Clam Pizza") 
	return c 
}

func (c *ClamPizza) Box() interface{}{
	fmt.Println("Boxing Clam Pizza") 
	return c 
}

func (k *KinshasaStore) CreatePizza(item string) (Pizza , error) {
	var pizza Pizza 

	switch item {
	case VEGGIE_PIZZA:
		pizza = new(VeggiePizza)
	case CLAM_PIZZA:
		pizza = new(ClamPizza)
	default: 
		err := errors.New("Unrecognized pizza type")
		return nil, err 
	}

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza, nil 
}

type GomaPizzaStore struct{} 

func (gS *GomaPizzaStore) CreatePizza(item string) (Pizza, error){
	var pizza Pizza 

	switch item{
	case VEGGIE_PIZZA: 
		pizza = new(VeggiePizza)
	case CLAM_PIZZA:
		pizza = new(ClamPizza) 
	default: 
		err := errors.New("Unrecognized pizza type") 
		return nil, err 
	}

	pizza.Prepare() 
	pizza.Bake() 
	pizza.Cut() 
	pizza.Box()

	return pizza, nil 
}

func main(){
	gomaPizzaStore := &GomaPizzaStore{} 
	gomaPizza, _:= gomaPizzaStore.CreatePizza(CLAM_PIZZA)
	kinshasaStore := &KinshasaStore{} 
	kinPizza, _ := kinshasaStore.CreatePizza(VEGGIE_PIZZA) 
	fmt.Println("Kinshasa pizza", kinPizza) 
	fmt.Println("Goma pizza", gomaPizza)
}
