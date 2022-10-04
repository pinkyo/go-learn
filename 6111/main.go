package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	fmt.Println(spiralMatrix(3, 3, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, nil}}}}}}}}}))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = -1
		}
	}
	cur := head
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	directionIdx := 0
	i, j := 0, 0
	for cur != nil {
		result[i][j] = cur.Val
		fmt.Printf("%v\n", result)
		idx := directionIdx % 4
		round := directionIdx / 4
		ii := i + directions[idx][0]
		jj := j + directions[idx][1]
		if ii < round || ii >= m-round || jj < round || jj >= n-round || (ii == round && jj == round) {
			directionIdx = directionIdx + 1
			idx := directionIdx % 4
			ii = i + directions[idx][0]
			jj = j + directions[idx][1]
		}
		i = ii
		j = jj
		cur = cur.Next
	}
	return result
}
