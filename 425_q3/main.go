package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(minArraySum([]int{2, 8, 3, 19, 3}, 3, 1, 1))
	// fmt.Println(minArraySum([]int{2, 4, 3}, 3, 2, 1))
	// fmt.Println(minArraySum([]int{3, 7, 1, 6}, 3, 2, 3))
	fmt.Println(minArraySum([]int{882, 307, 624, 469, 329, 684, 851, 608, 317, 205}, 431, 9, 4))
}

func minArraySum(nums []int, k int, op1 int, op2 int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	i := 0
	for ; op1 > 0; i++ {
		if op2 > 0 && (nums[i]+1)/2 < k {
			break
		}
		nums[i] = (nums[i] + 1) / 2
		if op2 > 0 {
			nums[i] -= k
			op2--
		}
		op1--
	}
	sub := nums[i:]
	sort.Sort(sort.Reverse(sort.IntSlice(sub)))
	for ii := len(sub) - 1; ii >= 0; ii-- {
		if op2 == 0 {
			break
		}
		if sub[ii] >= k {
			sub[ii] -= k
			op2--
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sub)))
	for ii := 0; ii < op1; ii++ {
		sub[ii] = (sub[ii] + 1) / 2
	}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}
