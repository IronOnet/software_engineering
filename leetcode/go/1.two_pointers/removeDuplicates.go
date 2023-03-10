// Given an array of sorted numbers, remove all duplicates from it. 
// you should not use any extra space; after removing the duplicates 
// in-place return the new length of the array 
package main 

import "fmt"

// Solution, In this problem we need to remove the duplicates in-place such 
// that the resultant length of the array remains sorted. As the input array 
// is sorted, therefore, one way to do this is to shift the elements left 
// whenever we encounter duplicates. In other words, we will keep one pointer 
// for iterating the array and one pointer for placing the next non-duplicate 
// number. So our algorithm will be to iterate the array and whenever we see 
// a non-duplicate number we move it next to the last non-duplciate number we've seen.
func RemoveDuplicates(nums []int) int{
	//left := 0 
	right := 1 
	for i:= 0; i < len(nums); i++{
		if nums[right-1] != nums[i]{
			nums[right] = nums[i] 
			right++
		}
	}
	return right
	
}

func RemoveDuplicates2(nums []int)int{
	nonDuplicateNext := 1
	for i := 0; i < len(nums); i++{
		if nums[nonDuplicateNext-1] == nums[i]{
			nums[nonDuplicateNext] = nums[i] 
			nonDuplicateNext++
		}
	}
	return nonDuplicateNext
}

func RemoveDuplicates3(nums []int)int{
	nextNonDuplicate := 1 
	for i := 0; i < len(nums); i++{
		if nums[nextNonDuplicate-1] == nums[i]{
			nums[nextNonDuplicate] = nums[i] 
			nextNonDuplicate++
		}
	}
	return nextNonDuplicate
}

func main(){
	array1 := []int{1, 1, 2, 2, 3, 3} 
	array2 := []int{1, 2, 4, 4, 5, 6, 7, 7, 9}
	fmt.Printf("The length of array %v after removing duplicates is %d\n", array1, RemoveDuplicates(array1))
	fmt.Printf("The length of array %v after removing duplicates is %d\n", array2, RemoveDuplicates(array2))
}