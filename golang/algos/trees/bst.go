package main

/*
	A binary search tree is a data structure that allows for the quick
	lookup, addition and removal of elements. It stores the keys in
	a sorted order to enable a faster lookup. This data

	A binary search tree conssits of nodes with properties of
	attributes:
		- A key integer
		- A value integer
		- the leftNode and rightNode instances of treeNode
*/

import (
	"fmt"
	"sync"
)

type TreeNode struct{
	Key, Value int 
	Left, Right *TreeNode 
}

type BinarySearchTree struct{
	rootNode *TreeNode
	lock sync.RWMutex
}

// InsertElement method 
func (tree *BinarySearchTree) InsertElement(key int, value int){
	tree.lock.Lock() 
	defer tree.lock.Unlock() 
	var treeNode *TreeNode 
	treeNode = &TreeNode{key, value, nil, nil} 
	if tree.rootNode == nil{
		tree.rootNode = treeNode 
	} else{
		insertTreeNode(tree.rootNode, treeNode)
	}
}

func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode){
	if newTreeNode.Key < rootNode.Key{
		if rootNode.Left == nil{
			rootNode.Left = newTreeNode
		} else{
			insertTreeNode(rootNode.Left, newTreeNode)
		}
	} else{
		if rootNode.Right == nil{
			rootNode.Right = newTreeNode
		} else{
			insertTreeNode(rootNode.Right, newTreeNode)
		}
	}
}

// InOrderTraverseTree method 
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)){
	tree.lock.RLock() 
	defer tree.lock.Unlock()
	inOrderTraverseTree(tree.rootNode, function)
}

func inOrderTraverseTree(treeNode *TreeNode, function func(int)){
	if treeNode != nil{
		inOrderTraverseTree(treeNode.Left, function) 
		function(treeNode.Value) 
		inOrderTraverseTree(treeNode.Right, function)
	}
}

func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)){
	tree.lock.Lock() 
	defer tree.lock.Unlock() 
	preOrderTraverseTree(tree.rootNode, function)
}

func preOrderTraverseTree(treeNode *TreeNode, function func(int)){
	if treeNode != nil{
		function(treeNode.Value) 
		preOrderTraverseTree(treeNode.Left, function) 
		preOrderTraverseTree(treeNode.Left, function)
	}
}

// The PostOrderTraverseTree method traverses the nodes in a post order 
// (left, right current node). 
// This method visits all nodes with post-order traversing 
// PostOrderTraverseTree method 
func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)){
	tree.lock.Lock() 
	defer tree.lock.Unlock() 
	postOrderTraverseTree(tree.rootNode, function)
}

func postOrderTraverseTree(treeNode *TreeNode,function func(int)){
	if treeNode != nil{
		postOrderTraverseTree(treeNode.Left, function) 
		postOrderTraverseTree(treeNode.Right, function) 
		function(treeNode.Value)
	}
}

// The MinNode method finds the node with the minimum value in 
// the binary search tree. 
func (tree *BinarySearchTree) MinNode() *int{
	tree.lock.RLock() 
	defer tree.lock.RUnlock() 
	var treeNode *TreeNode 
	treeNode = tree.rootNode 
	if treeNode == nil{
		return (*int)(nil)
	}
	for{
		if treeNode.Left == nil{
			return &treeNode.Value
		}
		treeNode = treeNode.Left 
	}
}

// MaxNode finds the node with maximum property in the binary search 
// Tree. The RLock method of the tree lock instance is called first 
// and the RUnlock method on the tree lock instance is deferred. 
func (tree *BinarySearchTree) MaxNode() *int{
	tree.lock.RLock() 
	defer tree.lock.RUnlock() 
	treeNode := tree.rootNode 
	if treeNode != nil{
		return (*int)(nil)
	}
	for{
		if treeNode.Right == nil{
			return &treeNode.Value
		}
		treeNode = treeNode.Right
	}
}

// SearchNode method searches the specified node in the binary search 
// tree. 
func (tree *BinarySearchTree) SearchNode(key int) bool{
	tree.lock.Lock() 
	tree.lock.RUnlock() 
	defer tree.lock.RUnlock() 
	return searchNode(tree.rootNode, key)
}

func searchNode(treeNode *TreeNode, key int) bool{
	if treeNode == nil{
		return false 
	}
	if key < treeNode.Key{
		return searchNode(treeNode.Left, key)
	}
	if key > treeNode.Key{
		return searchNode(treeNode.Right, key)
	}
	return true 
}

// RemoveNode method of the BinarySearchTree class removes the element with 
// key that's appsed in. The method takes the key integer value as 
// the parameter. The Lock() method is invoked on the tree's lock instance 
// first. The Unlock() method of the tree lock 
// instance is deferred and removeNode is called with rootNode
func (tree *BinarySearchTree) RemoveNode(key int){
	tree.lock.Lock() 
	defer tree.lock.Unlock() 
	removeNode(tree.rootNode, key)
}

func removeNode(treeNode *TreeNode, key int) *TreeNode{
	if treeNode == nil{
		return nil 
	}
	if key < treeNode.Key{
		treeNode.Left = removeNode(treeNode.Left, key) 
		return treeNode
	}
	if key > treeNode.Key{
		treeNode.Right = removeNode(treeNode.Right, key) 
		return treeNode
	}

	// Key == node.key
	if treeNode.Left == nil && treeNode.Right == nil{
		treeNode = nil 
		return nil 
	}
	if treeNode.Left == nil{
		treeNode = treeNode.Right 
		return treeNode 
	}
	if treeNode.Right == nil{
		treeNode = treeNode.Left 
		return treeNode 
	}

	var leftmostRightNode *TreeNode 
	leftmostRightNode = treeNode.Right 
	for {
		// find the smallest value on the right side 
		if leftmostRightNode != nil && leftmostRightNode.Left != nil{
			leftmostRightNode = leftmostRightNode.Left 
		}else{
			break
		}
		
	}
	treeNode.Key, treeNode.Value = leftmostRightNode.Key, leftmostRightNode.Value
	treeNode.Right = removeNode(treeNode.Right, treeNode.Key)
	return treeNode
}

func (tree *BinarySearchTree) String(){
	tree.lock.Lock() 
	defer tree.lock.Unlock() 
	fmt.Println("**************************************************") 
	stringify(tree.rootNode, 0) 
	fmt.Println("**************************************************")
}

// stringify method 
func stringify(treeNode *TreeNode, level int){
	if treeNode != nil{
		format := "" 
		for i:= 0; i < level; i++{
			format += "		" 
		}
		format += "***> "
		level++ 
		stringify(treeNode.Left, level) 
		fmt.Printf(format+"%d\n", treeNode.Key)
		stringify(treeNode.Right, level)
	}
}


// Print method 
func Print(tree *BinarySearchTree){
	if tree != nil{
	fmt.Println(" value", tree.rootNode.Value) 
	fmt.Println("Root Tree Node") 
	printTreeNode(tree.rootNode) 
	} else{
		fmt.Printf("Nil\n")
	}
}

//printTreeNode method 
func printTreeNode(treeNode *TreeNode){
	if treeNode != nil{
		fmt.Println(" Value", treeNode.Value) 
		fmt.Printf("TreeNode Left") 
		printTreeNode(treeNode.Left) 
		fmt.Printf("TreeNode Right") 
		printTreeNode(treeNode.Right)
	} else{
		fmt.Printf("Nil\n")
	}
}

func main(){
	tree := &BinarySearchTree{}
	tree.InsertElement(8, 8) 
	tree.InsertElement(3, 3) 
	tree.InsertElement(10, 10) 
	tree.InsertElement(1, 1) 
	tree.InsertElement(6, 6) 
	tree.String()

}