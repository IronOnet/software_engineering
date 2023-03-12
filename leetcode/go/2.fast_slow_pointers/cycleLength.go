// Problem statement, If a Linked list has a cycle
// Find the length of that cycle 
//
// Solution 
// We can use the solution to the "findiing cycles" problem to find the 
// length of the cyle. Once the fast and slow pointers meet, we can save 
// the slow pointer and iterate through the whole cycle with another 
// pointer untill we see the slow pointer again to find the length of the cycle. 
package main 

import "fmt" 

type ListNode struct{
	Val int 
	Next *ListNode 
} 

func HasCycle(head *ListNode) bool{
	if head == nil || head.Next == nil{
		return false 
	} 
	slowPtr := head 
	fastPtr := head 
	for slowPtr != nil && fastPtr.Next != nil{
		slowPtr = slowPtr.Next 
		fastPtr = fastPtr.Next.Next 
		if slowPtr == fastPtr{
			return true 
		}
	}
	return false 
} 

func CalcLengthCycle(head *ListNode) int{
	if !HasCycle(head){
		return 0 
	}
	slowPtr := head
	fastPtr := head 
	for slowPtr != nil && fastPtr.Next != nil{
		slowPtr = slowPtr.Next 
		fastPtr = fastPtr.Next 
		if slowPtr == fastPtr{
			return calcLengthCycle(slowPtr)
		}
	}
	return 0
	
}

func calcLengthCycle(slow *ListNode) int{
	current := slow 
	cycleLength := 0 
	for {
		current = current.Next 
		cycleLength++ 
		if !(current != slow){
			break
		}
	}
	return cycleLength 	
}


func main(){
	head := &ListNode{1, nil} 
	head.Next = &ListNode{2, nil}
	head.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next = &ListNode{5, nil} 
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}

	fmt.Printf("Cycle lenght for the linked list %v is : %d\n", CalcLengthCycle(head))
}
