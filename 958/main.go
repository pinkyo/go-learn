package main

import "fmt"

func main() {
	fmt.Println(isCompleteTree(&TreeNode{Val: 1, Left: &TreeNode{}, Right: &TreeNode{}}))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	visited := make([]bool, 101)
	var preOrder func(node *TreeNode, cur int)
	cnt := 0
	invalid := false
	preOrder = func(node *TreeNode, cur int) {
		if node == nil {
			return
		}
		cnt++
		if cur >= len(visited) {
			invalid = true
			return
		}
		visited[cur] = true
		preOrder(node.Left, cur*2+1)
		preOrder(node.Right, cur*2+2)
	}
	preOrder(root, 0)
	if invalid {
		return false
	}
	for i := 0; i < cnt; i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}
