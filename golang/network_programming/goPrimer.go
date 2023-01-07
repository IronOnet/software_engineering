package main 

import (
	"fmt"
	"os"
)

func slicesAndArrays(){


	var x [128]int 

	// creating a new slice with 
	// the new operator 
	xp := new([128]int)

	// Using the make operator

	xM := make([]int, 50, 100)
	xMPrime := new([100]int)[0:50] 


}

// Structures 
type Person struct{
	Name Name 
	Email []Email 
}


type Name struct{
	Family string 
	Personal string 
}

type Email struct{
	Kind string 
	Address string 
}


func structures(){
	person := Person{
		Name: Name{ Family: "Wanet", Personal: "Arnaud"}, 
		Email: []Email{Email { Kind: "home", Address: "andrewwanet9@gmail.com"},
						Email{Kind: "Work", Address: "arnaudwanetwork9@gmail.com"}}
	}
}


// Functions 
func checkError(err error){
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error()) 
		os.Exit(1)
	}
}

// methods 
func (p Person) String() string{
	s := p.Name.Personal + " " + p.Name.Family	
	for _, v := range p.Email{
		s += "\n" + v.Kind + ": " + v.Address
	}

	return s
}

// Type conversion 
func typeConversion(){
	var b []byte 
	b = []byte("string") 

	var s string 
	s = string(b[:])
}