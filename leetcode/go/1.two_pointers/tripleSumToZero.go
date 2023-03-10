// Given an array of unsorted numbers, find all niuqe triplets in it
// that add up to zero
// Example 1:
// Input: [-3, 0, 1, 2, -1, 1, -2]
// OUput: [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
// Explanation: there are four unique triplets whose sum is equal to Zero

package main

import (
	"fmt"
	"sort"
)


func FindTriplets(nums []int)[][]int{

	sort.Ints(nums) 

	var result[][]int 

	for i := 0; i < len(nums); i++{
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		searchPairs(nums, -nums[i], i + 1, &result);
	}
	return result
}

func searchPairs(nums []int, targetSum int, left int, triplets *[][]int){
	right := len(nums)-1; 
	for left < right{
		currentSum := nums[left] + nums[right] 
		if currentSum == targetSum{
			*triplets = append(*triplets, []int{-targetSum, nums[left], nums[right]})
			left++ 
			right--
			for left < right && nums[left] == nums[left - 1]{
				left++
			}
			for left < right && nums[right] == nums[right - 1]{
				right--
			}
		} else if (targetSum > currentSum){
			left++
		} else{
			right--
		}
	}

}

func main(){
	vector := []int{-3, 0, 1, 2, -1, 1, -2}
	result := FindTriplets(vector)
	vector2 := []int{-5, 2, -1, -2, 3} 
	result2 := FindTriplets(vector2)
	fmt.Printf("The triplets that add up to zero for the vector %v\n are: %v\n", vector, result)
	fmt.Printf("The triplets that add up to zero for the vector %v\n are: %v\n", vector2, result2)
}
