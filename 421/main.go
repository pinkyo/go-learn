package main

func main() {
}

type tier struct {
	left, right *tier
}

func findMaximumXOR(nums []int) int {
	root := &tier{}
	for _, nums := range nums {
		cur := root
		for i := 31; i >= 0; i-- {
			if nums&(1<<uint(i)) != 0 {
				if cur.right == nil {
					cur.right = &tier{}
				}
				cur = cur.right
			} else {
				if cur.left == nil {
					cur.left = &tier{}
				}
				cur = cur.left
			}
		}
	}
	return traverse(root, root, 31)
}

func traverse(first, second *tier, depth int) int {
	if depth < 0 {
		return 0
	}
	var result int
	if first.left != nil && second.right != nil {
		result = traverse(first.left, second.right, depth-1) + 1<<uint(depth)
	}
	if first.right != nil && second.left != nil {
		result = max(result, traverse(first.right, second.left, depth-1)+1<<uint(depth))
	}
	if result > 0 {
		return result
	}
	if first.left != nil && second.left != nil {
		result = traverse(first.left, second.left, depth-1)
	}
	if first.right != nil && second.right != nil {
		result = max(result, traverse(first.right, second.right, depth-1))
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
