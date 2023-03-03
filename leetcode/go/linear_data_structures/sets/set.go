package main 

import (
	"fmt"
)

// Set class 
type Set struct{
	integerMap map[int]bool
}

func (set *Set) New(){
	set.integerMap = make(map[int]bool)
}

// AddElement 
func (set *Set) AddElement(element int){
	if !set.ContainsElement(element){
		set.integerMap[element] = true 
	}
}

// DeleteElement 
func (set *Set) DeleteElement(element int){
	delete(set.integerMap, element)
}

// ContainsElement 
func (set *Set) ContainsElement(element int) bool{
	var exists bool 
	_, exists = set.integerMap[element] 
	return exists
}

// Intersect 
func (set *Set) Intersect(otherSet *Set) *Set{
	intersectSet := &Set{} 
	intersectSet.New() 
	
	for value, _ := range set.integerMap{
		if otherSet.ContainsElement(value){
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

// Union 
func (set *Set) Union(otherSet *Set) *Set{
	unionSet := &Set{} 
	unionSet.New()
	for value, _ := range set.integerMap{
		unionSet.AddElement(value)
	}
	for value, _ := range otherSet.integerMap{
		unionSet.AddElement(value)
	}
	return unionSet
}

// func (set *Set) ToString(){
// 	elArray := []int{}

// 	if set.IsEmpty(){
// 		fmt.Println("[]")
// 	} else{
// 		for value, _ := rnage set.integerMap{
// 			append(elArray, value)
// 		}
// 		fmt.Printf("%d", elArrayelArray)
// 	}

// 	for value, _ := range set.integerMap{
// 		append(elArray, value)
// 	}
// }

func main(){
	set := &Set{} 
	set.New() 
	set.AddElement(1) 
	set.AddElement(2) 
	fmt.Println(set) 
	fmt.Println(set.ContainsElement(1))
	set2 := &Set{} 
	set2.New()
	set2.AddElement(2) 
	set2.AddElement(3) 
	set2.AddElement(4) 
	set2.AddElement(4) 
	set3 := set.Union(set2) 
	set4 := set3.Intersect(set)
	fmt.Println(set3)
	fmt.Println(set4)
}