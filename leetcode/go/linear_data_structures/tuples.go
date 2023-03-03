// Tuples are finite ordered sequence of objects. They can 
// contain a mixture of other data types and are used 
// to group related data types into a data structure. 
// In a relational database a tuple is a row of a table

package main 

import (
	"fmt"
)

type Tuple struct{

}

// the h function "height" which returns the product of 
// parmeters x and y 
func h(x, y int) int{
	return x * y 
}

// the g function with returns x and y after modification 
func g(l, m int) (x, y int){
	x, y = 2*l, 4*m 
	return 
}

func main(){
	fmt.Print(h(g(3, 8)))
}



