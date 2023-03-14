// Problem statement 
// Given a list of intervals, merge all the overlapping intervals to produce a list 
// that has only mutually exclusive intervals. 
// Example 1: 
// Intervals: [[1,4], [2,5], [7,9]] 
// Output: [[1,5], [7,9]] 
// Explanation: since the first two intervals [1,4] and [2,5] overlap, we merged 
// them into one [1,5]

// Example 2 
// Intervals [[6,7], [2,4], [5,9]] 
// Output: [[2,4],[5,9]] 
// Explanation: since the intervals [6,7] and [5,9] overlap we merged them into 
// one [5,9].

// Solution 
// Let's take the example of two intervals (a and b) such that a.start <= b.start. 
// there are four possible scenarios. 
// 1. a and b do not overlap 
// 2. some part of b overlaps with a 
// 3. a fully overlaps b 
// 4. b fully overlaps a but both have the same start

// Our goal is to merge the intervals whenever they overlap. For the above mentioned 
// three overlapping scenarios (2, 3, and 4)
// When some part of b overlaps with a the resulting interval will be (a.start, b.end) 
// whhen a completely overlaps b the resulting interval is (a.start, a.end) 
// when b fully overlaps a but both have the same start the resulting interval will be (a.start, b.end)

// The algorithm will look like this
// merge the intervals at the start time to ensure a.start <= b.start 
// If 'a' overlaps 'b' (i.e b.start <= a.end), we need to merge them into 
// a new interval 'c' such that 
// c.start = a.start 
// c.end = max(a.end, b.end)

package main 

import "sort"

type Interval struct{
	Start int 
	End int 
}

type IArray []Interval



func MergeIntervals(intervals []Interval) []Interval{
	if len(intervals) < 2){
		return intervals
	} 

	// Sort Intervals by start time 
	// Have to write a custom sort function
	sort.Sort(intervals)
}

func overlap(intervalA, intervalB [2]int) bool{
	if intervalA[1] < intervalB[1] && (intervalA[0] < intervalB[0] || intervalA[0] > intervalB[0]){
		return true 
	}
}
