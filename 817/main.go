package main

func main() {
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	idxMap := make(map[int]int)
	cur := head
	idx := 0
	for cur != nil {
		idxMap[cur.Val] = idx
		idx++
		cur = cur.Next
	}
	flags := make([]bool, len(idxMap))
	for i := 0; i < len(nums); i++ {
		if v, ok := idxMap[nums[i]]; ok {
			flags[v] = true
		}
	}
	var result int
	var cnt int
	for i := 0; i < len(flags); i++ {
		if flags[i] {
			cnt++
			continue
		}

		if cnt > 0 {
			result++
		}
		cnt = 0
	}
	if cnt > 0 {
		result++
	}
	return result
}
