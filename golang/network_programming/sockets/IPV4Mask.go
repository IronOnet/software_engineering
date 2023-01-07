package main 

import (
	"fmt" 
	"net" 
	"os"
)

func main(){
	if len(os.Args) != 2{
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-address\n", os.Args[0]) 
		os.Exit(1)
	}

	dotAddr := os.Args[1] 

	addr := net.ParseIP(dotAddr) 
	if addr == nil{
		fmt.Println("Invalid Address") 
		os.Exit(1)
	}

	mask := addr.DefaultMask() 
	network := addr.Mask(mask) 
	ones, bits := mask.Size() 

	fmt.Println("Address is ", addr.String(), 
				"\nThe Bit length is ", bits, 
				"\nLeading Ones Count is", ones, 
				"\nMask is (hex) ", mask.String(), 
				"\nNetwork is : ,", network.String())
}