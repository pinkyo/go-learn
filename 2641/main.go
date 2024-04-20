package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	sumMap := make(map[int]int)
	preMap := make(map[*TreeNode]*TreeNode)

	var preorder func(node *TreeNode, depth int, fn func(node *TreeNode, depth int))
	preorder = func(node *TreeNode, depth int, fn func(node *TreeNode, depth int)) {
		if node == nil {
			return
		}
		fn(node, depth)
		if node.Left != nil {
			preorder(node.Left, depth+1, fn)
		}
		if node.Right != nil {
			preorder(node.Right, depth+1, fn)
		}
	}

	preorder(root, 0, func(node *TreeNode, depth int) {
		sumMap[depth] += node.Val
		if node.Left != nil {
			preMap[node.Left] = node
		}
		if node.Right != nil {
			preMap[node.Right] = node
		}
	})

	fmt.Println(sumMap)
	nextValueMap := make(map[*TreeNode]int)
	preorder(root, 0, func(node *TreeNode, depth int) {
		value := sumMap[depth]
		if preNode, ok := preMap[node]; ok {
			if preNode.Left != nil {
				value -= preNode.Left.Val
			}
			if preNode.Right != nil {
				value -= preNode.Right.Val
			}
		} else {
			value -= node.Val
		}
		nextValueMap[node] = value
	})

	for node, value := range nextValueMap {
		node.Val = value
	}
	return root
}
