package main

import "fmt"

/** KEY WORDS
	break
	case
	chan
	const
	continue
	default
	defer
	else
	fallthrough
	for
	func
	go
	goto
	if
	import
	interface
	map
	package
	range
	return
	select
	struct
	switch
	type
	var

	Constants:

		- true, false, iota, nil

	Types:
		 int, int8, int16, int32, int64
		 uint, uint8, uint16, uint32, uint64 uintptr
		 float32, float64, complex128 complex64
		 bool byte, rune string error

	Functions:
		 make, len, cap, new, append, copy, close, delete
		 complex real, imag panic recover
**/

func makeLoop(){
	fmt.Println("For loop without iterator initialization")
	i := 1
	for i <= 10{
		fmt.Println(i)
		i = i + 1
	}

	//return i
}

func forLooopWithInit(){
	fmt.Println("For loop with inline initialization")
	for i:= 0; i <= 10; i++{
		fmt.Println(i)
	}
}

func makeConditions(){
	for i:= 1; i <= 10; i++{
		if i % 2 == 0{
			fmt.Println(i, "even")
		} else{
			fmt.Println(i, "odd")
		}
	}
}

func ifElseControls(){
	i := 0
	if i == 0{
		fmt.Println("Zero")
	} else if i == 1{
		fmt.Println("One")
	} else if i == 2{
		fmt.Println("Two")
	} else if i == 3{
		fmt.Println("Three")
	}
}

func switchControls(){
	i := 0
	switch i{
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	default: fmt.Println("Uknown number")
	}
}

func main(){
	makeLoop()
	forLooopWithInit()
	makeConditions()
}
