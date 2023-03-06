// Adelson, Velski, and Landis pioneered the AVL tree 
// data structure and hence it is named after them. 
// it consists of height adjusting binary search trees. 
// the balance factor is obtained by finding the difference 
// between the heights of the left and right sub-trees. 
// balancing is done using rotation techniques. If the 
// balance factor is greater than one, rotation shifts the 
// nodes to the opposite of the left or right subtrees
package main

// KeyValue interface has the LessThan and EqualTo methods. 
// the LessThan and EqualTo methods take KeyValue as 
// parameter and return a Boolean value after checking 
// the less than or equal to condition. 

// KeyValue type 
type KeyValue interface{
	LessThan(KeyValue) bool 
	EqualTo(KeyValue) bool 
}

// TreeNode Class 
// this class has KeyValue, BalanceValue, and LinkedNodes as properties. 
// The AVL tree is created as tree of nodes of the TreeNode type 
type TreeNode struct{
	KeyValue KeyValue 
	BalanceValue int 
	LinkedNodes [2]*TreeNode
}