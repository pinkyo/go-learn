package main

func main() {
	// fmt.Println(btreeGameWinningMove())
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var postOrder func(node *TreeNode) int
	var leftSum int
	var rightSum int
	postOrder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := postOrder(node.Left)
		right := postOrder(node.Right)
		if node.Val == x {
			leftSum = left
			rightSum = right
		}
		return left + right + 1
	}
	sum := postOrder(root)
	thirdSum := sum - leftSum - rightSum - 1
	if leftSum > sum-leftSum || rightSum > sum-rightSum || thirdSum > sum-thirdSum {
		return true
	}
	return false
}
