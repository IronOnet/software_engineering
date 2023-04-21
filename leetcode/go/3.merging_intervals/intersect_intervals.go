package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

func IntersectIntervals(intervals_a, intervals_b []Interval) []Interval {
	if len(intervals_a) == 0 && len(intervals_b) == 0 {
		return intervals_a
	} else if len(intervals_a) == 0 {
		return intervals_b
	} else if len(intervals_b) == 0 {
		return intervals_a
	}

	result := []Interval{}
	i, j := 0, 0

	for i < len(intervals_a) && j < len(intervals_b) {
		// Check if the interval arr1[i] intersect with arr2[j]
		// check if one of the interval's start time lies with other intervals
		if intervals_a[i].Start >= intervals_b[j].Start && intervals_a[i].Start <= intervals_b[j].End ||
			intervals_b[j].Start >= intervals_a[i].Start && intervals_b[j].Start <= intervals_a[i].End {
			result = append(result, Interval{Max(intervals_a[i].Start, intervals_b[j].Start), Min(intervals_a[i].End, intervals_b[j].End)})
		}

		// Move next from the interval with finishing first
		if intervals_a[i].End < intervals_b[j].End {
			i++
		} else {
			j++
		}
	}

	return result

	// When does two interval intersect?
}

func IntersectIntervals2(arr1, arr2 []Interval) []Interval{

	result := []Interval{} 
	i, j := 0, 0 
	for i < len(arr1) && j < len(arr2){
		if arr1[i].Start >= arr2[j].Start && arr1[i].Start <= arr2[j].End || arr2[j].Start >= arr1[i].Start && arr2[j].End <= arr1[i].End{
			result = append(result, Interval{Max(arr1[i].Start, arr2[j].Start), Min(arr1[i].End, arr2[j].End)})
		}

		if arr1[i].End < arr2[j].End{
			i++
		} else {
			j++
		}
	}
	return result 
}


func IntersectIntervals3(arr1, arr2 []Interval) []Interval{
	result := []Interval{} 
	i, j := 0, 0 
	for i < len(arr1) && j < len(arr2){
		if arr1[i].Start >= arr2[j].Start && arr1[i].Start <= arr2[j].End || 
			arr2[j].Start >= arr1[i].Start && arr2[j].Start <= arr1[i].End{
				result = append(result, Interval{Max(arr1[i].Start, arr2[j].Start), Min(arr1[i].End, arr2[j].End)})
		}

		if arr1[i].End < arr2[j].End{
			i++
		} else{
			j++
		}
	}
	return result 
}


func IntersectIntervals4(intervals_a, intervals_b []Interval) []Interval{
	result := []Interval{}
	i, j := 0, 0  
	for i < len(intervals_a) && j < len(intervals_b){
		if intervals_a[i].Start >= intervals_b[j].Start && intervals_a[i].Start <= intervals_b[j].End || 
			intervals_b[j].Start >= intervals_a[i].Start && intervals_b[j].Start <= intervals_a[i].End{
				result = append(result, Interval{Max(intervals_a[i].Start, intervals_b[j].Start), Min(intervals_a[i].End, intervals_b[j].End)})
		}

		if intervals_a[i].End < intervals_b[j].End{
			i++ 
		} else{
			j++
		}
	}

	return result 
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input1 := []Interval{{1, 3}, {5, 6}, {7, 9}}
	input2 := []Interval{{2, 3}, {5, 7}}
	result := IntersectIntervals4(input1, input2)
	fmt.Printf("Intersection of intervals %v and intervals %v yields =====> %v\n", input1, input2, result)
}