package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func CanAttend(intervals []Interval) bool {
	if len(intervals) == 0{
		return false 
	}

	sort.SliceStable(intervals, func(i, j int) bool{
		return intervals[i].Start <= intervals[j].Start
	})

	for i := 1; i < len(intervals); i++{
		if intervals[i].Start < intervals[i-1].End{
			return false 
		}
	}
	return true 
}

func main() {
	appointments := []Interval{{6, 7}, {2, 4}, {8, 12}} // => should be true
	appointments2 := []Interval{{4, 5}, {2, 3}, {3, 6}} // => should be false
	fmt.Printf("Given appointments %v can patient attend all of them ===> %v\n", appointments, CanAttend(appointments)) 
	fmt.Printf("Given appointments %v can patient attend all of them ===> %v\n", appointments2, CanAttend(appointments2)) 
}