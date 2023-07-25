/**
	Problem statement

	We are given an array containing n objects. Each object, wehn created was
	assigned a number from 1 to n based on their current creation sequence.
	This means that the object with sequence number 3 was created just before
	the object with sequence number 4.

	Write a function to sort the objects in-place on their creation sequence number
	in O(n) and without extra space. For simplicity, let's assume we are passed an integer
	array containing only the sequence numbers, though each number is actually an
	object.

	Example 1:
	Input: [3, 1, 5, 4, 2]
	Output: [1, 2, 3, 4, 5]


	Example 2:
	Input [2, 6, 3, 1, 5]
	Output: [1, 2, 3, 4, 5]


	Solution:

	We iterate the array one number at a time, and if the current number
	We are iterating is not at the correct index, we swap it with the number
	at its correct index. This way we will go through all numbers and place
	them in their correct indices, hence, sorting the whole array.
**/

package main

import "fmt"


func SortCyclic(nums *[]int){
	i := 0 
	for i < len(*nums){
		j := (*nums)[i] - 1 
		if ((*nums)[i] != (*nums)[j]){
			(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		} else{
			i++
		}
	}
}

func main(){
	vec := []int{3, 1, 5, 4, 2} 
	vec2 := make([]int, len(vec))
	copy(vec2, vec)
	SortCyclic(&vec) 
	fmt.Printf("Before cyclick sort ===> %v\n After cyclic sort ===> %v\n", vec2, vec)
}