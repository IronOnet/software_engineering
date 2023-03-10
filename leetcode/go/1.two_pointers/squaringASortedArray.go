/*
Given a sorted array, create a new array containing squares of all the member
of the input array in the sorted order

Example 1:

Input: [-2, -1, 0, 2, 3]
Output: [0, 1, 4, 4, 9]

Example 2:
Input: [-3, -1, 0, 1, 2]
Output: [0, 1, 1, 4, 9]

Solution:

This question might seem straightforward. The only trick is that we can
have negative numubers in the input array, which will make it a bit difficult
to generate theoutput array with squares in sorted order.

An easier approach could be to first find the index of the first non-negative number
in the array. aFTER that we can use two pointers to iterate the array.

One pointer will move forward to iterate the non-negative numbers and the other
pointer will move backward to iterate the negative nubmers. At any step
whichever number gives us a bigger square will be added to the output array.

Since the numbers at both ends can give us the largest square, an alternate approach
could be to use two pointers starting at both the ends of the input array. At
any step, whichever gives us the bigger square we add it to the result
array and move to the next/previous number according to the pointer
*/
package main

import "fmt" 

func SquareArray(nums []int)[]int{

	resultArray := make([]int, len(nums)) 
	leftPointer, rightPointer := 0, len(nums)-1
	fmt.Println("Right pointer", rightPointer)
	fmt.Println("Result array length: ", len(resultArray)) 
	highestSquareIdx := len(nums)-1

	for leftPointer <= rightPointer{
		leftSquare := nums[leftPointer] * nums[leftPointer] 
		rightSquare := nums[rightPointer] * nums[rightPointer] 
		if (leftSquare > rightSquare){
			
			resultArray[highestSquareIdx] = leftSquare
			leftPointer++
		} else{
			resultArray[highestSquareIdx] = rightSquare 
			rightPointer--
		}
		highestSquareIdx--
	}
	return resultArray
}

func SquareArray2(nums []int) []int{
	resultArray := make([]int, len(nums)) 
	leftPointer, rightPointer := 0, len(nums)-1 

	highestSquareIndex := len(nums)-1 

	for leftPointer <= rightPointer{
		leftSquare := nums[leftPointer] * nums[leftPointer] 
		rightSquare := nums[rightPointer] * nums[rightPointer] 

		if leftSquare > rightSquare{
			resultArray[highestSquareIndex] = leftSquare
			leftPointer++
		}else{
			resultArray[highestSquareIndex] = rightSquare 
			rightPointer--
		}
		highestSquareIndex--
	}
	return resultArray
}


func main(){
	array1 := []int{-2, -1, 0, 2, 3}; 
	result := SquareArray2(array1) 
	fmt.Printf("Squared array for array %v is %v\n", array1, result)
}
