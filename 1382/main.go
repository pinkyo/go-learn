package main

import "fmt"

func main() {
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	inorder := func(root *TreeNode) []int {
		var result []int
		var helper func(root *TreeNode)
		helper = func(root *TreeNode) {
			if root == nil {
				return
			}
			helper(root.Left)
			result = append(result, root.Val)
			helper(root.Right)
		}
		helper(root)
		return result
	}
	arr := inorder(root)
	fmt.Println(arr)
	var helper func([]int) *TreeNode
	helper = func(value []int) *TreeNode {
		if len(value) == 0 {
			return nil
		}
		mid := len(value) / 2
		return &TreeNode{
			Val:   value[mid],
			Left:  helper(value[:mid]),
			Right: helper(value[mid+1:]),
		}
	}
	return helper(arr)
}
