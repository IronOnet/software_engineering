// Stack is a last in, first out structure in which items are added 
// from the top. Stacks are used in parsers for solving maze algorithms 
// Push, pop, top and get size are typical operations that are allowed 
// on stack data structures. 
package main  

import (
	"fmt" 
	"strconv" 
)

// Element class 
type Element struct{
	elementValue int 
}

// String method on ELement class 
func (el *Element) String() string{
	return strconv.Itoa(el.elementValue);
}

type Stack struct{
	elements []*Element 
	elementCount int 
}

// The New method 
// the New method on the Stack class creats a dyamic array of 
// elements. The Stack class has the count and array pointer 
// of elements. The code snippet with the Stack class 
// definition and the New method is as follows 

// NewStack returns a new stack 
func (stack *Stack) New(){
	stack.elements = make([]*Element, 0)
}

// The Push method 
// THe Push method adds the node to the top of the stack class 
// Push adds a node to the stack 
func (stack *Stack) Push(element *Element){
	stack.elements = append(stack.elements[:stack.elementCount], element) 
	stack.elementCount += 1
}

// Pop remvoes and returns a node from the stack 
func (stack *Stack) Pop() *Element{
	if stack.elementCount == 0{
		return nil 
	}
	var length int = len(stack.elements) 
	var element *Element = stack.elements[length-1] 
	
	if length > 1{
		stack.elements = stack.elements[:length-1]
	} else{
		stack.elements = stack.elements[0:]
	}
	stack.elementCount = len(stack.elements) 
	return element 
}

func main(){
	var stack *Stack = &Stack{} 
	stack.New() 
	
	element1 := &Element{3} 
	element2 := &Element{5} 
	element3 := &Element{7} 
	element4 := &Element{9} 

	stack.Push(element1) 
	stack.Push(element2) 
	stack.Push(element3) 
	stack.Push(element4) 

	fmt.Println(stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop(), stack.Pop())
}