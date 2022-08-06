package main

import "fmt"

// If the root node is nil, then the tree is univalued, otherwise, if the root node's value is not
// equal to the left or right node's value, then the tree is not univalued, otherwise, the tree is
// univalued if the left and right subtrees are univalued
func main() {
	fmt.Println(isUnivalTree(&TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}))
}

// Definition for a binary tree node.
// A TreeNode is a node in a binary tree.
// @property {int} Val - the value of the node
// @property Left - The left child of the node.
// @property Right - The right subtree of a node contains only nodes with keys greater than the node's
// key.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// If the root node is nil, return true; otherwise, if the root node's value is not equal to the value
// of the root node, return false; otherwise, return the result of recursively calling the function on
// the left and right subtrees
func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	var preOrder func(*TreeNode) bool
	preOrder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		// Checking if the value of the node is not equal to the value of the root node.
		if node.Val != val {
			return false
		}
		// Calling the function recursively on the left and right subtrees.
		return preOrder(node.Left) && preOrder(node.Right)
	}
	return preOrder(root)
}
