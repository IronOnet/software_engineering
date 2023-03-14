// Problem
// Find the middle node of a linked list 
// Given the head of a Singly LinkedList, write a method to return 
// the middle node of the LinkedList. 
// If the total number of nodes in the LinkedList is even, return the second middle node 

// Solution 
// Using a brute force strategy, we can count the numbers of nodes N
// in the first iteration and in the second iteration 
// stop at node N//2 and return it, this approach will have a time complexity 
// of O(N+N) and a space complexity of O(1) 

// A better approach would be to use the fast and slow pointer method. 
// we use two pointers slow and fast, we will increment the fast 
// pointer twice more than the slow pointer, the slow pointer will always 
// return the middle of the linked list 



package main 


type LinkedNode struct{
	Val int 
	Next *LinkedNode 
} 

func FindMiddle(head *LinkedNode) *LinkedNode{
	slow, fast := head, head 
	for slow != nil && fast.Next != nil{
		slow = slow.Next 
		fast = fast.Next.Next 
	} 
	return slow 
}
