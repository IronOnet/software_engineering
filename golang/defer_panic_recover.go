package main 

import (
	"fmt"
	"os"
)

// Go has a special statement called defer which schedules a function 
// call to be run after the function completes

func first(){
	fmt.Println("first")
}

func second(){
	fmt.Println("second")
}

// Basically defer moves the call to second to the end 
// of the function
// it is often used when resources need to be freed in 
// some way.

func openFile(filename string){
	f, _ := os.Open(filename)
	defer f.Close()
}

// Panic and recover 
// the panic function causes a runtime error 
// we can handle a runtime panic with the 
// builtin recover function 
func panicRecover(){
	// here the call to recover will 
	// never execute because the panic 
	// function immediately stops the execution 
	// of the function
	panic("PANIC")
	str := recover()
	fmt.Println(str)
}

func panicRecoverDefer(){
	defer func(){
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
func main(){
	defer first() 
	second()
	//defer panicRecover()
	panicRecoverDefer()
}