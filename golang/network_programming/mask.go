/* 
	The IP mask type 
	An IP address is typically divided into the components of 
	the subnetwork, the subnet, the address and the device
*/
package main 

import (
	"fmt" 
	"net" 
	"os"
	"strconv"
)
// In order to handle masking operations 
// you need to use this type 
// type IPMask []byte 

func main(){
	if len(os.Args) != 4{
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-address ones bits\n", os.Args[0])
		os.Exit(1)
	}

	dotAddr := os.Args[1] 
	ones, _ := strconv.Atoi(os.Args[2])
	bits, _ := strconv.Atoi(os.Args[3])

	addr := net.ParseIP(dotAddr) 
	if addr == nil{
		fmt.Println("Address is invalid") 
		os.Exit(1)
	}

	mask := net.CIDRMask(ones, bits)
	network := addr.Mask(mask)
	fmt.Println("Adress is ", addr.String(), 
				"\nThe mask length is: ", bits, 
				"\nThe Leading Ones count is: ", ones, 
				"\nThe Mask is (hex): ", mask.String(), 
				"\nThe Network is: ", network.String()) 
	os.Exit(0)
}