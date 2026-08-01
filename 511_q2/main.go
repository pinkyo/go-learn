package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countDominantNodes(root *TreeNode) int {
	cnt := 0
	var postOrder func(node *TreeNode) int
	postOrder = func(node *TreeNode) int {
		maxVal := node.Val
		if node.Left != nil {
			maxVal = max(maxVal, postOrder(node.Left))
		}
		if node.Right != nil {
			maxVal = max(maxVal, postOrder(node.Right))
		}
		if node.Val >= maxVal {
			cnt++
		}
		return maxVal
	}

	postOrder(root)
	return cnt
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
