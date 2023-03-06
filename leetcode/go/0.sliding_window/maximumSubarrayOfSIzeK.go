// Given an array of positive numbers and a positive number K
// find the maximum sum of any contiguous subarray of size K
package main 

import "fmt"


// Time complexity O(N) 
// Space Complexity O(1)
func MaximumSubarrayOfSizeK(nums []int, K int) int{
	windowStart := 0 
	windowSum := 0 
	maxValue := 0 

	for windowEnd := 0; windowEnd < len(nums); windowEnd++{
		windowSum += nums[windowStart] 
		if windowEnd >= K - 1{
			maxValue = Max(maxValue, windowSum)
			windowSum -= nums[windowStart] 
			windowStart++
		}
	}
	return maxValue
}

func MaximumSubarrayOfSizeK2(nums []int, K int) int{
	windowStart := 0 
	maxValue := 0 
	windowSum := 0 

	for windowEnd := 0; windowEnd < len(nums); windowEnd++{
		windowSum += nums[windowStart] 
		if windowEnd > K - 1{
			maxValue = Max(maxValue, windowSum)
			windowSum -= nums[windowStart] 
			windowStart++ 
		}
	}
	return maxValue
}

func Max(a, b int) int{
	if a > b{
		return a
	}
	return b
}

func main(){
	k := 3 
	array := []int{2, 1, 5, 1, 3, 2}
	value := MaximumSubarrayOfSizeK(array, k)
	fmt.Printf("The maximum value of any contiguous subarray of size %d in %v is %d", k, array, value)
}