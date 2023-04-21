package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

type Queue []*Interval 


func (q *Queue) Add(interval *Interval){
	if len(*q) == 0{
		*q = append(*q, interval)
	} else{
		appended := false 
		//var addedInterval *Interval 
		for i, addedInterval := range *q{
			if interval.Start > addedInterval.Start{
				*q = append((*q)[:i], append(Queue{interval}, (*q)[i:]...)...)
				appended = true 
				break 
			}
		}
		if !appended{
			*q = append(*q, interval)
		}
	}
}



func MinimumMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	result := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			result += 1
		}
	}

	return result
}

func main() {
	meeting1 := []Interval{{1, 4}, {2, 5}, {7, 9}}         // Should be 2
	meeting2 := []Interval{{6, 7}, {2, 4}, {8, 12}}        // should be 1
	meeting3 := []Interval{{1, 4}, {2, 3}, {3, 6}}         // should be 2
	meeting4 := []Interval{{4, 5}, {2, 3}, {2, 4}, {3, 5}} // should be 2

	fmt.Printf("Minimum # rooms to hold meetings %v is ====> %d\n", meeting1, MinimumMeetingRooms(meeting1)) // 2
	fmt.Printf("Minimum # rooms to hold meetings %v is ====> %d\n", meeting2, MinimumMeetingRooms(meeting2)) // 1 
	fmt.Printf("Minimum # rooms to hold meetings %v is ====> %d\n", meeting3, MinimumMeetingRooms(meeting3)) // 2
	fmt.Printf("Minimum # rooms to hold meetings %v is ====> %d\n", meeting4, MinimumMeetingRooms(meeting4)) // 2
}