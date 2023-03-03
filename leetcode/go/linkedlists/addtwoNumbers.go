package main

import "fmt"

/**
Add Two Numbers.
You are given two non-negative integers. The digits are stored in
reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

YOu may assume the two numbers do not contain any leading zero,
except the number 0 itself.

Inputs: l1 = [2, 3, 4], l2 = [5, 6, 4]
Output: [7, 0, 8]
Explanation: 342 + 465 = 807

INputs: l1= [0], l2 = [0]
Output: [0]

INput: l1 = [9, 9, 9, 9, 9, 9, 9], l2 = [9, 9, 9, 1]
Output: [8, 9, 9, 9, 0, 0, 0, 1]

COnstraints:

- The number of nodes in each linked list is in the range [1, 100]
*/

type ListNode struct{
	Val int 
	Next *ListNode 
}

type LinkedList struct{
	Head *ListNode 
	//Tail *ListNode 
}

type ListBuilder struct{
	l *LinkedList
}

func (b *ListBuilder) AddNode(val int) *ListBuilder{
	b.l.AddNode(val) 
	return b 
}

func (b *ListBuilder) GetList() *LinkedList{
	return b.l
}

func (l *LinkedList) AddNode(val int){
	node := new(ListNode) 
	node.Val = val 
	if node.Next == nil{
		node.Next = l.Head
	}
	l.Head = node
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode){
	dummy := new(ListNode) 
	l3 := &dummy 
	var carry int 
	list1, list2 := l1, l2  

	for (list1 != nil) || (list2 != nil){
		l3 = &((*l3).Next) 
		*l3 = new(ListNode) 

		var sum int 

		if list1 != nil{
			sum += list1.Val 
			list1 = list1.Next 
		}

		if list2 != nil{
			sum += list2.Val 
			list2 = list2.Next 
		}

		(*l3).Val = (sum + carry) % 10 
		carry = (sum + carry) / 10 
	}

	if carry > 0{
		l3 = &((*l3).Next) 
		*l3 = new(ListNode) 
		(*l3).Val = carry
	}
}


func main(){
	myList := new(LinkedList) 
	myList.AddNode(3) 
	myList.AddNode(4) 
	myList.AddNode(5)
	fmt.Println((*myList))
}