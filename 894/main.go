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

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return nil
	}
	if n == 1 {
		return []*TreeNode{{}}
	}
	var result []*TreeNode
	for i := 1; i < n; i += 2 {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - i - 1)
		for _, l := range left {
			for _, r := range right {
				result = append(result, &TreeNode{
					Val:   0,
					Left:  l,
					Right: r,
				})
			}
		}
	}
	return result
}
