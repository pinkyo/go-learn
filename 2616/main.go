package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
	fmt.Println(minimizeMax([]int{4, 2, 1, 2}, 1))
}

func minimizeMax(nums []int, p int) int {
	if p == 0 {
		return 0
	}
	sort.Ints(nums)
	l, h := 0, nums[len(nums)-1]-nums[0]
	for l <= h {
		mid := (l + h) / 2
		if check(nums, mid, p) {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func check(nums []int, mid, p int) bool {
	var cnt int
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] <= mid {
			cnt++
			i++
		}
		if cnt >= p {
			return true
		}
	}
	return false
}
