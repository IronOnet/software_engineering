// A Qeue consist of elements to be processed
// In a particular order or based on priority 
// A priority-bnased queu of orders is shown in the 
// following code snippet structured as a heap. 

// linear data structures and a sequential collection. 
package main 

import (
	"fmt"
)

// Qeue - Array of Orders type
type Queue []*Order  

// Order class 
type Order struct{
	priority int 
	quantity int 
	product string 
	customerName string 
}

// The New method on the Order class assigns the properties from 
func (order *Order) New(priority int, quantity int, product string, customerName string){
	order.priority = priority 
	order.quantity = quantity
	order.product = product 
	order.customerName = customerName 
}

// Add method adds the order to the queue 
func (queue *Queue) Add(order *Order){
	if len(*queue) == 0{
		*queue = append(*queue, order)
	} else{
		appended := false 
		for i, addedOrder := range *queue{
			if order.priority > addedOrder.priority{
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true 
				break
			}
		}
		if !appended{
			*queue = append(*queue, order)
		}
	}
}

func main(){
	queue := make(Queue, 0) 
	order1 := &Order{} 
	priority1 := 2 
	quantity1 := 20 
	var product1 string = "Computer" 
	var customerName1 string = "Greg White" 
	order1.New(priority1, quantity1, product1, customerName1) 

	var order2 *Order = new(Order) 
	var priority2 int = 1 
	var quantity2 int = 10 

	var product2 string = "Monitor" 
	var customerName2 string = "John smith" 
	order2.New(priority2, quantity2, product2, customerName2) 
	queue.Add(order1) 
	queue.Add(order2) 
	

	for i:= 0; i < len(queue); i++{
		fmt.Println(queue[i])
	}


}