package main

import "sort"

type Interval struct {
	Start int
	End   int
}

func InsertInterval(intervals []Interval, interval Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	intervals = append(intervals, interval)

	sort.SliceStable(intervals, func(i, j int) bool{
		return intervals[i].Start <= intervals[j].Start 
	})

	merged := []Interval{intervals[0]} 
	for i:= 1; i < len(intervals); i++{
		lastInterval := &merged[len(merged)-1] 
		if intervals[i].Start <= lastInterval.End{
			if lastInterval.End < intervals[i].End{
				lastInterval.End = intervals[i].End 
			}
		} else{
			merged = append(merged, intervals[i])
		}
	}

	return merged
}