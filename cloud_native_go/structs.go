package main 

type Person struct{
	name string 
	age int
}

func main(){
	var p = Person{} 
	var p2 = Person{"bob", 21} 
	var p3 = { name: "Bob", age: 21 } 
	var p4 = &Person{} 
	var p5 = &Person{ "bob", 21 }  
	var p6 = &Person{ "bob", 21 }
}