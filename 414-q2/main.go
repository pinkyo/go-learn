package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(maxPossibleScore([]int{1000000000, 0}, 1000000000))
	// fmt.Println(maxPossibleScore([]int{100, 1000000000, 0}, 1009))
	fmt.Println(maxPossibleScore([]int{1000000000, 1000, 0}, 1000))
}

func maxPossibleScore(start []int, d int) int {
	sort.Ints(start)
	l, r := 0, (start[len(start)-1]-start[0]+d)/(len(start)-1)
	if l == r {
		return l
	}
	for l <= r {
		mid := (l + r + 1) >> 1
		if check(start, d, mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l - 1
}

func check(start []int, d, mid int) bool {
	cur := start[0]
	for i := 1; i < len(start); i++ {
		if start[i]+d-cur < mid {
			return false
		}
		if cur+mid < start[i] {
			cur = start[i]
		} else {
			cur = cur + mid
		}
	}
	return true
}
