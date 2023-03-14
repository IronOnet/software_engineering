// Given the head of a Singly linkedList that contains a cyle, write a function to find the 
// starting node of the cycle. 

// Solution 
// Using the fast and slow pointer approach, we can find the start of the cylcle 
// by taking two pointers, pointer1 & pointer 2 
// initializing both pointers to the start of the linkedlist 
// We find the length of the cycle let's call it K 
// Move pointer2 ahead by K nodes 
// Now keepe incrementing pointer1 and pointer2 until they both meet 
// As pointer2 is K nodes ahead of pointer1, which means pointer2 must have 
// completed one loop in the cycle when both pointers meet. Their meeting 
// point will be the start of the cycle. 

package main 

type LinkedNode struct{
	Value int 
	Next *LinkedNode 
}

func FindStartCycle(head *LinkedNode) *LinkedNode{
	fast, slow := head, head
	cycleLength := 0  
	for fast != nil && fast.Next != nil{
		slow = slow.Next 
		fast = fast.Next.Next 
		if slow == fast{
			cycleLength = calcCycleLength(slow)			
			break 
		}	
	}

	return findStart(head, cycleLength)

}

func calcCycleLength(slow *ListNode) int{
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

func findStart(head *ListNode, cycleLength int) *ListNode{
	pointer1 := head 
	pointer2 := head 

	for cycleLength > 0{
		pointer2 = pointer2.Next 
		cycleLength-- 
	} 

	for pointer1 != pointer2{
		pointer1 = pointer1.Next 
		pointer2 = pointer2.Next 
	} 
	return pointer1 
}
