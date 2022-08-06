package main

import (
	"fmt"
)

func main() {
	// fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 2))
	fmt.Println("Hello World!")
}

// If k is 0, then we count the number of elements that appear more than once; otherwise, we count the
// number of elements that are k apart
func findPairs(nums []int, k int) int {
	set := make(map[int]int)
	for _, v := range nums {
		set[v]++
	}
	var result int
	for kk, v := range set {
		if k == 0 {
			if v > 1 {
				result++
			}
			continue
		}
		if _, ok := set[kk+k]; ok {
			result++
		}
	}

	return result
}
