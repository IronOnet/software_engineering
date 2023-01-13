package main 

import (
	"fmt" 
	"os"
)

func main(){
	var goos string = os.Getenv("GOOS") 
	fmt.Printf("The Operating System is %s\n", goos) 
	path := os.Getenv("PATH") 
	fmt.Printf("The System Path is %s\n", path)
}
