package main

import (
	"fmt"
	"sort"
)

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthLargestPerfectSubtree finds the k-th largest perfect subtree in a binary tree.
// A perfect subtree is a subtree where the sum of all node values is equal for the left and right subtrees.
// The function uses post-order traversal to calculate the sums and stores them in a slice.
// It then sorts the slice in descending order and returns the k-th largest sum.
// If there are fewer than k perfect subtrees, it returns -1.
func kthLargestPerfectSubtree(root *TreeNode, k int) int {
	result := make([]int, 0)

	var postOrderTraverse func(*TreeNode) (bool, int)
	postOrderTraverse = func(node *TreeNode) (bool, int) {
		if node == nil {
			return false, 0
		}

		if node.Left == nil && node.Right == nil {
			result = append(result, 1)
			return true, 1
		}

		left, leftSum := postOrderTraverse(node.Left)
		right, rightSum := postOrderTraverse(node.Right)

		if left && right && leftSum == rightSum {
			result = append(result, leftSum+rightSum+1)
		}

		return left && right && leftSum == rightSum, leftSum + rightSum + 1
	}

	postOrderTraverse(root)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	fmt.Printf("%v\n", result)

	if len(result) < k {
		return -1
	}
	return result[k-1]
}
