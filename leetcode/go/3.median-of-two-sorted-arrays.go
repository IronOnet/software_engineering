package main

import "math"

func findMedianOfTwoSortedArrays(nums1, nums2 []int) float64 {
	A, B := nums1, nums2
	total := len(nums1) + len(nums2)
	half := (total + 1) / 2

	var aLeft, aRight float64
	var bLeft, bRight float64

	if len(B) < len(A) {
		A, B = B, A
	}

	left, right := 0, len(A)-1
	for {
		i := (left + right) >> 1 //A
		j := half - i - 2        // B

		if i >= 0 {
			aLeft = float64(A[i])
		} else {
			aLeft = math.Inf(-1)
		}

		if (i+1) < len(A){
			aRight = float64(A[i+1])
		} else{
			aRight = math.Inf(1)
		}

		if j>= 0{
			bLeft = float64(B[j]) 
		} else{
			bLeft = math.Inf(-1)
		}

		if (j+1) < len(B){
			bRight = float64(B[j+1])
		} else{
			bRight = math.Inf(1)
		}

		// partition is correct 
		if aLeft <= bRight && bLeft <= aRight{
			// Odd 
			if total % 2 == 1{
				return max(aLeft, bLeft)
			}
			// even 
			return (max(aLeft, bLeft) + min(aRight, bRight)) / 2
		} else if aLeft > bRight{
			right = i - 1 
		} else{
			left = i + 1
		}
	}
}

func max(a, b float64) float64{
	return math.Max(a, b)
}

func min(a, b float64) float64{
	return math.Min(a, b)
}