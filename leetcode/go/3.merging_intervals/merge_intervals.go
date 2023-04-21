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

import (
	"fmt"
	"sort"
)


type Interval struct{
	Start int 
	End int 
}

func MergeIntervals(intervals []Interval) []Interval{
	if len(intervals) == 0{
		return intervals 
	}

	sort.SliceStable(intervals, func(i, j int) bool{
		return intervals[i].Start <= intervals[j].Start
	})

	merged := []Interval{intervals[0]} 

	for i := 1; i < len(intervals); i++{
		last := &merged[len(merged)-1]
		if intervals[i].Start <= last.End{
			if last.End < intervals[i].End{
				last.End = intervals[i].End 
			}
		}else{
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

func InsertInterval(interval Interval, intervals []Interval) []Interval{
	if len(intervals) == 0{
		return intervals
	}
	
	intervals = append(intervals, interval) 

	sort.SliceStable(intervals, func(i, j int) bool{
		return intervals[i].Start <= intervals[j].Start 
	})

	merged := []Interval{intervals[0]} 

	for i := 1; i < len(intervals); i++{
		last := &merged[len(merged)-1] 
		if intervals[i].Start <= last.End{
			if last.End < intervals[i].End{
				last.End = intervals[i].End 
			}
		} else{
			merged = append(merged, intervals[i])
		}
	}

	return merged 
}


func MergeIntervals2(intervals []Interval) []Interval{
	if len(intervals) == 0{
		return intervals
	}

	sort.SliceStable(intervals, func(i, j int) bool{
		return intervals[i].Start <= intervals[j].Start 
	})

	merged := []Interval{intervals[0]} 

	for i := 1; i < len(intervals); i++{
		lastInterval := &merged[len(merged)-1] 
		if intervals[i].Start <= lastInterval.End{
			if lastInterval.End < intervals[i].End{
				lastInterval.End = intervals[i].End 
			}
		} else{
			intervals = append(intervals, intervals[i])
		}
	}

	return intervals
}

func main(){
	intervals := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals2 := []Interval{{1, 4}, {2, 5}, {7, 9}}  
	intervals3 := []Interval{{6, 7}, {2, 4}, {5, 9}} 
	intervals4 := []Interval{{1, 4}, {2, 6}, {3, 5}}
	merged2 := MergeIntervals(intervals2) 
	merged3 := MergeIntervals(intervals3) 
	merged4 := MergeIntervals(intervals4)
	merged := MergeIntervals(intervals) 
	fmt.Printf("interval %v merged ===> %v\n", intervals,  merged)
	fmt.Printf("interval %v merged ===> %v\n", intervals2, merged2)
	fmt.Printf("interval %v merged ===> %v\n", intervals3, merged3)
	fmt.Printf("interval %v merged ===> %v\n", intervals4, merged4)
}