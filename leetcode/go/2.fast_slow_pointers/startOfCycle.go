// Problem statement 
// Given the head of a Singly LinkedList that contains a cycle, write a function 
// to find the starting node of the given cycle.


// Solution : 
// We first need to find if our linked list contains a cycle
// 1. We take two pointers. let's call them pointer1 and pointer2 
// 2. Initialize both pointers to point to the start of the LinkedList 
// 3. We can find the length of the LinkedList cycle using the approach discussed in 
// the "LinkedList Cycle" problem. Let's assume that the length of the cycle is K nodes 
// 4. Move pointer2 ahead by K nodes 
// 5. Now, keep incrementing pointer1 and pointer2 until they both meet. 
// 6. As pointer2 is 'K' nodes ahhead of pointer1, which means, pointer2 must have completed one 
// loop in the cycle when both pointers meet. Their meeting point will be the start of the cycle.
package main 

import "fmt"

type ListNode struct{
	Val int 
	Next *ListNode
}

func FindCycleStart(head *ListNode) *ListNode{
	cycleLength := 0 
	slow := head
	fast := head 
	for slow != nil && fast.Next != nil{
		fast = fast.Next.Next 
		slow = slow.Next 
		if (slow == fast){
			cycleLength = calcCycleLength(slow);
			break
		}	
	}

	return findStart(head, cycleLength)
}

func calcCycleLength(slow *ListNode) int {
	current := slow 
	cycleLength := 0 
	for{
		current = current.Next 
		cycleLength++
		if !(current != slow){
			break 
		} 
	} 
	return cycleLength
}

func findStart(head *ListNode, cycleLength int) *ListNode{
	pointer1 := head 
	pointer2 := head 

	for cycleLength > 0{
		pointer2 = pointer2.Next 
		cycleLength--; 
	} 

	// Increment both pointers until they meet at the start of the cycle 
	for pointer1 != pointer2{
		pointer1 = pointer1.Next 
		pointer2 = pointer2.Next 
	} 

	return pointer1
}


func main(){
	head := &ListNode{1, nil} 
	head.Next = &ListNode{2, nil} 
	head.Next.Next = &ListNode{3, nil} 
	head.Next.Next.Next = &ListNode{4, nil} 
	head.Next.Next.Next.Next = &ListNode{5, nil} 
	head.Next.Next.Next.Next.Next = &ListNode{6, nil}

	head.Next.Next.Next.Next.Next.Next = head.Next.Next; 
	fmt.Printf("The cycle for the list %v starts at %v\n", head, FindCycleStart(head)) 

	head.Next.Next.Next.Next.Next = head.Next.Next.Next 
	fmt.Println("The cycle for the list %v starts at %v\n", head, FindCycleStart(head)) 

	head.Next.Next.Next.Next.Next.Next = head 
	fmt.Println("The cycle for the list %v starts at %v\n", head, FindCycleStart(head)) 
}
