package main

import (
	"fmt"
	"sync"
)

type TreeNode struct{
	key int 
	value int 
	leftNode *TreeNode 
	rightNode *TreeNode
}

// BinarySearchTree class 
type BinarySearchTree struct{
	rootNode *TreeNode
	lock sync.RWMutex
}

// The InsertElement method 
func (tree *BinarySearchTree) InsertElement(key int, value int){
	tree.lock.Lock() 
	defer tree.lock.Unlock() 
	var treeNode *TreeNode 
	treeNode = &TreeNode{key, value, nil, nil} 
	if tree.rootNode == nil{
		tree.rootNode = treeNode 
	}else{
		insertTreeNode(tree.rootNode, treeNode)
	}
}

// insertTreeNode recursively insert nodes in the tree
func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode){
	if newTreeNode.key < rootNode.key{
		if rootNode.leftNode == nil{
			rootNode.leftNode = newTreeNode
		} else{
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else{
		if rootNode.rightNode == nil{
			rootNode.rightNode = newTreeNode 
		} else{
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

// inOrderTraverse method visits all nodes in order. the RLock() method 
// on the tree lock instance is called first. The RUnlock() method is 
// deferred on the tree lock instance before invoking the inOrderTraverseTree 
// method. 
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)){
	tree.lock.RLock() 
	defer tree.lock.RUnlock() 
	inOrderTraverseTree(tree.rootNode, function)
}

// The inOrderTraverseTree method traverses the left, the root, and the right 
// tree. The inOrderTraverseTree method takes treeNode of the TreeNode type 
// and function as parameters. The inOrderTraverseTree method is called 
// on the leftNode and rightNOde with function as a parameter 

func inOrderTraverseTree(treeNode *TreeNode, function func(int)){
	if treeNode != nil{
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

// The PreOrderTraverseTree method 
// The PreOrderTraverseTree method visits all tree nodes with preorder 
// traversing. the tree lock instance is locked first and
func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)){
	tree.lock.Lock() 
	defer tree.lock.Unlock() 
	preOrderTraverseTree(tree.rootNode, function)
}

// preOrderTraverseTree method 
func preOrderTraverseTree(treeNode *TreeNode, function func(int)){
	if treeNode != nil{
		function(treeNode.value) 
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

// PostOrderTraverseTree method 
// This metho traverses the nodes in a post order (left, righht, current node) 
//
func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)){
	tree.lock.Lock() 
	defer tree.lock.Unlock() 
	postOrderTraverseTree(tree.rootNode, function)
}

// postOrderTraverseTree method 
func postOrderTraverseTree(treeNode *TreeNode, function func(int)){
	if treeNode != nil{
		postOrderTraverseTree(treeNode.leftNode, function) 
		postOrderTraverseTree(treeNode.rightNode, function) 
		function(treeNode.value)
	}
}

// The MinNode method 
// MinNode finds the node with the minimum value in the binary search tree. 
// In the following code snippet, the RLock method of the tree lock instance 
// is invoked first and the RUnlock method on the tree lock isntance 
// is deferred 
func (tree *BinarySearchTree) MinNode() *int{
	tree.lock.RLock() 
	defer tree.lock.RUnlock() 
	treeNode := tree.rootNode 
	if treeNode == nil{
		return (*int)(nil)
	}
	for {
		if treeNode.leftNode == nil{
			return &treeNode.value 
		}
		treeNode = treeNode.leftNode
	}
}

// MaxNode finds the node with the maximum value in the 
// binary search tree
func (tree *BinarySearchTree) MaxNode() *int{
	tree.lock.RLock() 
	defer tree.lock.RUnlock()
	treeNode := tree.rootNode 
	if treeNode == nil{
		return (*int)(nil)
	}
	for{
		if treeNode.rightNode == nil{
			return &treeNode.value 
		}
		treeNode = treeNode.rightNode
	}
}

// SearchNode 
// this method searches the specified node in the binary search tree. 
func (tree *BinarySearchTree) SearchNode(key int) bool{
	tree.lock.RLock() 
	defer tree.lock.RUnlock() 
	return searchNode(tree.rootNode, key)
}

// searchNode takes treeNode, a pointer of TreeNode type 
// and akey integer value as parameters. 
func searchNode(treeNode *TreeNode, key int) bool{
	if treeNode == nil{
		return false 
	}
	if key < treeNode.key{
		return searchNode(treeNode.leftNode, key)
	}
	if key > treeNode.key{
		return searchNode(treeNode.rightNode, key)
	}
	return true 
}

// The RemoveNOde method 
// the REMOVEnODE method of the BinarySearchTree class remvoes the 
// element with key that's passed in. The method takes the key integer 
// value as the parameter. the Lock() method is invoked on the 
// tree's lock instance first. The Unlock() method of the tree lock 
// instance is deferred, and removeNOde is called with rootNode 
// and the key value as parameters 
func (tree *BinarySearchTree) RemoveNode(key int){
	tree.lock.Lock() 
	defer tree.lock.Unlock() 
	removeNode(tree.rootNode, key)
}

// removeNode method 
// this method takes treeNode fo the TreeNode type and 
// a key integer value as parameters. 
func removeNode(treeNode *TreeNode, key int) *TreeNode{
	if treeNode == nil{
		return nil
	}
	if key < treeNode.key{
		treeNode.leftNode = removeNode(treeNode.leftNode, key) 
		return treeNode 
	}
	if key > treeNode.key{
		treeNode.rightNode = removeNode(treeNode.rightNode, key) 
		return treeNode 
	}
	// key === node.key 
	if treeNode.leftNode == nil && treeNode.rightNode == nil{
		treeNode = nil 
		return nil 
	}
	if treeNode.leftNode == nil{
		treeNode = treeNode.rightNode
		return treeNode 
	}
	if treeNode.rightNode == nil{
		treeNode = treeNode.leftNode 
		return treeNode 
	}

	var leftmostRightNode *TreeNode 
	leftmostRightNode = treeNode.rightNode
	for{
		if leftmostRightNode != nil && leftmostRightNode.leftNode != nil{
			leftmostRightNode = leftmostRightNode.leftNode 
		} else{
			break 
		}
	}
	treeNode.key, treeNode.value = leftmostRightNode.key, leftmostRightNode.value 
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.key) 
	return treeNode
}

func (tree *BinarySearchTree) String(){
	tree.lock.Lock() 
	defer tree.lock.Unlock() 
	fmt.Println("-------------------------------------------") 
	stringify(tree.rootNode, 0) 
	fmt.Println("-------------------------------------------")
}

func stringify(treeNode *TreeNode, level int){
	if treeNode != nil{
		format := "" 
		for i:= 0; i < level; i++{
			format += " " 

		}
		format += "---[ " 
		level++
		stringify(treeNode.leftNode, level) 
		fmt.Printf(format + "%d\n", treeNode.key) 
		stringify(treeNode.rightNode, level)
	}
}

func main(){
	var tree *BinarySearchTree = &BinarySearchTree{} 
	tree.InsertElement(8, 8) 
	tree.InsertElement(3, 3) 
	tree.InsertElement(10, 10) 
	tree.InsertElement(1, 1) 
	tree.InsertElement(6, 6) 
	tree.String()
}