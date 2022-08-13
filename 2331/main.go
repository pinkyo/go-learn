package main

func main() {
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		switch root.Val {
		case 0:
			return false
		case 1:
			return true
		}
		return false
	}
	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)
	switch root.Val {
	case 2:
		return left || right
	case 3:
		return left && right
	}
	return false
}
