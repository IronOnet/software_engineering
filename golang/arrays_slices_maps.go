package main 

import "fmt" 

// Arrays 
func makeArray(){
	// an array is a numbered sequence of elements 
	// of a single type with a fixed length 
	var arrayX [5]int // this is an array composed of 5 ints
	arrayX[4] = 100 
	fmt.Println(arrayX)


}

func floatArrays(){
	var x[5]float64 
	x[0] = 98 
	x[1] = 93
	x[2] = 77 
	x[3] = 82 
	x[4] = 83

	var total float64 = 0 
	for i:= 0; i < 5; i++{
		total += x[i]
	}
	fmt.Println(total/5)
}

func loopTypeConversion(){
	var x[5] float64 
	x[4] = 100 
	var total float64 = 0 
	for _, value := range x{
		total += value 
	}
	fmt.Println(total / float64(len(x)))
}

func loopTypeConversionUpdated(){
	var x[5] float64 
	x[4] = 100 
	var total float64 = 0 
	// since the index is not used in the loop 
	// we can omit it 
	for _, value := range x{
		total += value 
	}
	fmt.Println(total/float64(len(x)))
}

// Shorter syntax for creating arrays 
func arraysShorter(){
	x := [5]float64{98, 93, 77, 82, 83} 

	// you can also break it up like this 
	x2 := [5]float64{
		98, 
		93, 
		77, 
		82, 
		83,
	}

	fmt.Println(x, x2)
}

// Slices  
// A slice is a segment of an array 
// like arrays slices are indexable 
// and have a length. Unlike arrays 
// this length is allowed to change
func slicesDemo(){
	// here's an example of a slice 
	var x[]float64 

	// you can use the builtin make() function 
	// to create a slice 
	// the following line creates a slice 
	// that is associated with an uderlying 
	// float64 array of length 5
	// slices are always associated with some 
	// array, although they can never be longer 
	// than the array, they can be smaller
	x2 := make([]float64, 5)

	// there is another way of creating slices 
	// in go using the [high:low] operator 
	arr := []float64{1, 2, 3, 4, 5} 
	x3 := arr[0:5]

	fmt.Println(x, x2, x3)
}


// Slice functions 
// go includes two buil-in functions to assist with 
// slices: append and copy
func sliceFunctions(){
	slice1 := []int{1, 2, 3} 
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func copySliceDemo(){
	slice1 := []int{1, 2, 3} 
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

/* 
	Maps 
	A map is an unordered collection of key-value pairs
	Also known as an associative array, a hash table 
	// or a dictionary
*/

func mapsDemo(){
	// the map type is represented by the keyword 
	// map, followed by the key type in brackets and 
	// finaly the value type
	//var x map[string]int

	// like arrays and slices maps can be accessed 
	// using brackets

	// The following code returns a runtime error 
	// runtime errors happen when you run the program 
	// while compile-time errors happen at compilation timee

	// maps have to be initialized before 
	// they can be used
	//x["key"] = 10 

	// we should have written this 
	x := make(map[string]int)
	x["key"] = 10 
	fmt.Println(x["key"])
	fmt.Println(x)

	// we can access the length of a map with 
	// the len function 
	// and can remove items from it with the delete function 

}


func mapDemo2(){
	elements := make(map[string]string)
	elements["H"] = "Hydrogen" 
	elements["He"] = "Helium" 
	elements["Li"] = "Lithium" 
	elements["Be"] = "Berylium" 
	elements["B"] = "Boron"
	elements["C"] = "Carbon" 
	elements["N"] = "Nitrogen" 
	elements["O"] = "Oxygen" 
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Ne"]) 
	fmt.Println(elements)

	// technicall a map returns the zero value
	// for the value type (which for strings is the empty string)
	// although we could check the zero value
	// in a condition (elements["Un"] == "") Go provides 
	// a better way 
	name, ok := elements["Un"]
	fmt.Println(name, ok)

	// Accessing an element of a map can return two values
	// instead of just one. The first value is the result 
	// of the lookup and the second value tells whether or not 
	// the lookup was successful.

	// like with arrays there is also a shorter way to 
	// create maps 
	// 
	elements2 := map[string]string{
		"H": "Hydrogen", 
		"He": "Helium", 
		"Li": "Lithium", 
		"Be": "Berylium", 
		"B": "Boron", 
	}

	fmt.Println(elements2)

	// maps are also used to store general information 
	elements3 := maps[string]maps[string]string{
		"H": map[string]string{
			"name": "Helium", 
			"state": "gas",
		}, 
		"He": map[string]string{
			"name": "Helium", 
			"state": "gas"
		}, 
		"Li": map[string]string{
			"name": "Lithium", 
			"state": "solid",
		},
	}

	fmt.Println(elements3) 

	if el, ok := elements3["Li"]; ok{
		fmt.Println(el["name"], el["state"])
	}

}

func main(){
	// makeArray() 
	// floatArrays()
	// sliceFunctions()
	// copySliceDemo()
	//mapsDemo() 
	mapDemo2()
}

