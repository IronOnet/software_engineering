// Given an array of positive numbers and a positive number "S", find the
// length of the smallest contigous subarray whose sum is greater than or
// equal to "S". Return 0, if no such subarray exists.
// Example:
// Input: [2, 1, 5, 2, 3, 2], S=7
// Output: 2
// Explanation: The smallest subarray with a sum great than or equal to '7' is [5, 2]
package main 

import "fmt"
import "math"


func MinSubarrayWithSum(nums []int, S int) int{
	windowStart := 0 
	windowSum := 0 
	minLength := math.MaxInt

	for windowEnd := 0; windowEnd < len(nums); windowEnd++{
		windowSum += nums[windowEnd] 
		for windowSum >= S  {
			minLength = Min(minLength, windowEnd - windowStart + 1) 
			windowSum -= nums[windowStart] 
			windowStart++	
		}//windowStart++
	}

	if minLength < 0{
		return 0
	}

	
	return minLength
}

func MinSubarrayWithSum2(nums []int, S int) int{
	windowStart := 0 
	minLength := math.MaxInt 
	windowSum := 0 

	for windowEnd := 0; windowEnd < len(nums); windowEnd++{
		windowSum += nums[windowEnd] 
		for windowSum >= S{
			minLength = Min(minLength, windowEnd - windowStart + 1)
			windowSum -= nums[windowStart] 
			windowStart++
		}
	}

	if minLength < 0{
		return 0 
	}

	return minLength
}

func MinSubarrayWithSum3(nums []int, S int) int{
	minLength := math.MaxInt 
	windowStart := 0 
	windowSum := 0 

	for windowEnd := 0; windowEnd < len(nums); windowEnd++{
		windowSum += nums[windowEnd] 
		for windowSum >= S{
			minLength = Min(minLength, windowEnd - windowStart + 1) 
			windowSum -= nums[windowStart] 
			windowStart++
		}
	}

	return minLength
}

func Min(a, b int) int{
	if a < b{
		return a 
	}
	return b 
}

func main(){
	array1 := []int{2, 1, 5, 2, 3, 2} 
	s1 := 7 
	array2 := []int{2, 1, 5, 2, 8} 
	s2 := 7 
	array3 := []int{3, 4, 1, 1, 6} 
	s3 := 8 
	result1, result2, result3 := MinSubarrayWithSum(array1, s1), MinSubarrayWithSum(array2, s2), MinSubarrayWithSum(array3, s3)
	
	fmt.Printf("The length of the smallest subarray whose sum is %d for array %v is : %d\n", s1, array1, result1)
	fmt.Printf("The length of the smallest subarray whose sum is %d for array %v is : %d\n", s2, array2, result2)
	fmt.Printf("The length of the smallest subarray whose sum is %d for array %v is : %d\n", s3, array3, result3)
}