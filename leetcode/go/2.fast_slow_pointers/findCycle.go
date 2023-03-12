// Given the head of a Singly Linkedlist, write a functino to determine if the LinkedList has a cycle 
// in it or not.

// Solution 
// Imagine two racers running in a circular racing track. If one racer is faster than the other, the faster 
// racer is bound to catch up and cross the slower racer from behind. We can use this fact to devise 
// an algorithm to determine if a LinkedList has a cycle int or not. 

// Imagine we have a slow and fast pointer to traverse the LinkedList. In each iteration, the slow pointer 
// moves one step and the fast pointer moves two steps. This gives us two conclusions: 

// 1. If the LinkedList doesn'thave a cycle in it, the fast pointer will reach the end of the LinkedList before 
// the slow pointer to reveal that there is no cycle in the LinkedList 
// 2. The slow pointer will never be able to catch up to the fast pointer if there is no cycle in the 
// LinkedList
package main 

import "fmt"

type ListNode struct{
	Val int 
	Next *ListNode
}

func HasCycle(node *ListNode) bool {
	if node == nil || node.Next == nil{
		return false 
	}	
	fastPointer := node 
	slowPointer := node 
	for slowPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next 
		slowPointer = slowPointer.Next 
		if slowPointer == fastPointer{
			return true 
		}
	}
	return false 
		
}

func HasCycle2(node *ListNode) bool{
	if node == nil || node.Next == nil{
		return false 
	} 
	slowPtr := head 
	fastPtr := head 
	for slowPtr != nil && fastPtr.Next != nil{
		fastPtr = fastPtr.Next.Next 
		slowPtr = slowPtr.Next 
		if slowPtr == fastPtr{
			return true 
		}
	} 
	return false
}


func main(){
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, nil} 
	head.Next.Next = &ListNode{3, nil} 
	head.Next.Next.Next = &ListNode{4, nil} 
	head.Next.Next.Next.Next = &ListNode{5, nil} 
	head.Next.Next.Next.Next  = &ListNode{6, nil}
	boolVal := HasCycle(head)

	fmt.Printf("Does the linked list %v has cycles? : %v\n", *head, boolVal)

	head.Next.Next.Next.Next.Next.Next = head.Next.Next 
	fmt.Printf("Does the linked list %v has cycles? : %v\n", *head, HasCycle(head))
}
