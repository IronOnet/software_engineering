package main 

import "fmt" 

func main(){
	var x = "Hello World" 
	var x2 string 
	x2 = "Hello Again!!"
	fmt.Println(x) 
	fmt.Println(x2)

	// this is how you declare a constant 
	const xy string = "Hello Planet Venus" 
	fmt.Println(xy)

	// Defining multiple variables 
	// use the keyword var or const 
	// followed by parentheses with each 
	// variable on its own line
	var (
		a = 5 
		b = 10 
		c  = 15
	)
}


