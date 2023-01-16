package main 

import "fmt" 

// the average function will need to take in a slice 
// of float64s and return one float64. 
func average(xs []float64) float64{
	//panic("Not Implemented") 
	total := 0.0 
	for _, v := range xs{
		total += v 
	}

	return total / float64(len(xs))
}


// functions are built up in a stack 
// Each time we call a function we push it
// into the call stack and each time we return
// from a function we pop the last function off of 
// the stack.

// we can name the return types of functions in go 
func f1() (r int){
	r = 1 
	return 
}

// returnning multiple values 
func f2() (int, int){
	return 5, 6
}

// multiple values are often used to return an error 
// value along with the result (x, err:= f()), or a 
// boolean to indicate success (x, ok := f())

// Variadic functions 
// There is a special form available for the last parameter 
// in a Go function

// By using three dots ... before the type name of the last 
// parameter you can indicate that it takes zero or more 
// of these parameters. In the following case we take 
// zero or more ints.
func add(args ...int)int{
	total := 0
	for _, v := range args{
		total += v
	}

	return total 
}

// we can also pass a slice of ints 
func variadicSlice(){
	xs := []int{1, 2, 3, 4, 5} 
	fmt.Println(add(xs...))
}

// Closures 
// It is possible to create functions 
// inside of functions 

// When you create functions like this 
// it has also access to other local variables
func closure(){
	add := func(x, y int) int{
		return x + y 
	}

	fmt.Println(add(6, 7))
}

func incrementClosure(){
	x := 0 
	increment := func() int{
		x++ 
		return x
	}
	fmt.Println(increment()) 
	fmt.Println(increment())
}

// One way to use a closure is by writing a function 
// which returns another function which - when called
// can generate a sequence of numbers 
func makeEvenGenerator() func() uint{
	i := uint(0)
	return func() (ret uint){
		ret = i 
		i += 2 
		return 
	}
}

func main(){
	x := []float64{30, 25, 20, 50, 60} 
	fmt.Println(average(x)) 

	// calling a function with multiple 
	// return values 
	y, z := f2()
	fmt.Println(y, z)
	fmt.Println("Using the variadic functiona add: ", add(1, 2, 3)) 
	fmt.Println("Printing a variadic slice ") 
	variadicSlice()
	fmt.Println("Closure examples") 
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) 
	fmt.Println(nextEven()) 
	fmt.Println(nextEven())
}

