// Given an array, find the average of all contigous subarrays
// of size 'K' in it.
package main

import "fmt"

// The efficient way to solve this problem would be to visualize
// each contigous subarray as a sliding window of 5 elements
// this means that when we move on to the next subarry, we will
// slide the window by one element. so to reuse the sum from the
// previsou subarray, we will subtract the element going out of the
// window and add the element now being included in the sliding window.
// this will save us from going through the whole subarray to find the
// sum and, as a result, the algorithm complexity will reduce to O(N)
func AverageOfSubarrayK(nums []int, K int) []float64{
	result := make([]float64, len(nums) - K + 1)
	var windowSum float64 = 0 
	var windowStart int = 0 
	
	for windowEnd := 0; windowEnd < len(nums); windowEnd++{
		windowSum += float64(nums[windowEnd])
		if windowEnd >= K - 1{
			result[windowStart] = windowSum / float64(K) 
			windowSum -= float64(nums[windowStart])
			windowStart++
		}
	}

	return result 
}

func AverageOfSubarrayK2(nums []int, K int) []float64{
	result := make([]float64, len(nums) - K + 1) 
	windowStart := 0 
	windowSum := 0.0 

	for windowEnd := 0; windowEnd < len(nums); windowEnd++{
		windowSum += float64(nums[windowEnd]) 
		if windowEnd >= K - 1{
			result[windowStart] = windowSum / float64(K) 
			windowSum -= float64(nums[windowStart]) 
			windowStart++
		}
	}
	return result
}
func main(){
	nums := []int{1, 3, 2, 6, -1, 4, 1, 8, 2} 
	K:= 5
	result := AverageOfSubarrayK(nums, K)
	result2 := AverageOfSubarrayK2(nums, K)
	fmt.Printf("Averages of subarrays of size %d of the array %v is %v\n", K, nums, result)
	fmt.Printf("Averages of subarrays of size %d of the array %v is %v\n", K, nums, result2)
}