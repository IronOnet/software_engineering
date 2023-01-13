// Basics of the Go programming language 
package main  

// calling the imports first 
import (
	"fmt"
	"os"
)

// Defining multiple types using 
// a compressed representation
type (
	IZ int 
	Fr string 
	ITA string 
)


const c = "C" 

var v int = 5 

type T struct{

}

// The init function initializes 
// a package 
func init(){

	var a int 
	Func1()
	fmt.Println(a)
}


func (t T) Method1(){

}

func Func1(){

}


// constants 
const PI = 3.14159

// Multiple assignment of constants is allowed 
// in Go 
const (
	Monday, Thuesday, Wednesday = 1, 2, 3 
	Thursday, Friday, Saturday = 4, 5, 6
)

// Constants can be used for enumeration 
const (
	Unknown = 0 
	Female = 1 
	Male = 2
)

// The value iota can also be used to enumerate 
// values. The first use of iota is set to 0 
// and the remaining ones are incremented by 1
const (
	a = iota 
	b = iota 
	c = iota 
)

// This can be shortened to 
const (
	a1 = iota 
	b1 
	c1
)


// You can give enumeration a type name 
type Color int 

const (
	RED Color = iota 
	ORANGE 
	YELLOW 
	GREEN 
	BLUE 
	VIOLET 
	INDIGO
)


// Variables are declared using the key word var 
var somVariable int 

// variables can also be declared as pointers 
var a2, b2 *int // two variable pointers of type int 

// Declaration and assignment can be combined in the general format 
var a3 int = 15 
var i = 5 
var b bool = false 
var str string = "Go says hello world"


// Variables could also be expressions computed 
// at runtime 
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

// Value types and reference types 
// In Go pointers are reference types, as well as slices, maps and channels 
// The variables that are referenced are stored in the Heap, which is 
// garbage collected and which as a much larger memory stack than the 
// stack. 

// The init() function 

var PI2 float64 

func init(){
	PI2 = 4 * math.Atan(1) // init function computes PI
}