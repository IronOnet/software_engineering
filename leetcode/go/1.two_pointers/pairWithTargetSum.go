// Given an array of sorted numbers and a target sum, find a pair in the
// array whose sum is equal to the given target.
// Write a function to return the indices of the indices of the
// two numbers. (i.e the pair) such that they add up to the given target
package main 

import "fmt"

func pairWithTargetSum(nums []int, T int) []int{
	left := 0 
	right := len(nums) - 1
	for left < right{
		sum := nums[left] + nums[right] 
		if sum == T{
			return []int{left, right}
		} else if sum < T{
			left++
		} else{
			right--
		}
	}
	return []int{-1, -1}
}

func main(){
	target1, target2 := 6 , 11
	array1, array2 := []int{1, 2, 3, 4, 5, 6}, []int{2, 5, 9, 11}
	fmt.Printf("The pair with the target sum of %d for the array %v is %v\n", target1, array1, pairWithTargetSum(array1, target1))
	fmt.Printf("The pair with the target sum of %d for the array %v is %v\n", target2, array2, pairWithTargetSum(array2, target2))

}