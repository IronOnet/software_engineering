package main

// A tree is a non-linear data structure. Trees are used for 
// search and other use cases. A binary tree has nodes 
// that have a maximum of two children. A binary search 
// tree consists of nodes where the property values 
// of the left node are less than the property values 
// of the right node

// A binary search tree is a data structure that allows for the quick 
// lookup, addition and removal of elements. It stores the keys 
// in a sorted order to enable a faster lookup. 
// On average, space usage for a binary search tree is 
// of the order of O(n), whereas the insert, search and 
// delete operations are of the order O(log n). 
// A binary search tree consists of nodes 
// with properties or attributes 
// A key integer 
// A value integer 
// The leftNode and rightNode instances of a TreeNode 

// TreeNode class 
type TreeNode struct{
	key int 
	value int 
	leftNode *TreeNode 
	rightNode *TreeNode
}

