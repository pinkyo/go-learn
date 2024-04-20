package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(minimumSeconds([]int{2, 1, 3, 3, 2}))
	fmt.Println(minimumSeconds([]int{8, 13, 3, 3}))
}

func minimumSeconds(nums []int) int {
	firstMap := make(map[int]int)
	lastMap := make(map[int]int)
	maxMap := make(map[int]int)
	for i, e := range nums {
		if pre, ok := lastMap[e]; ok {
			val := (i - pre) / 2
			if maxMap[e] < val {
				maxMap[e] = val
			}
		}
		if _, ok := firstMap[e]; !ok {
			firstMap[e] = i
		}
		lastMap[e] = i
	}
	for k, last := range lastMap {
		first := firstMap[k]
		val := (first + len(nums) - last) / 2
		if maxMap[k] < val {
			maxMap[k] = val
		}
	}

	result := math.MaxInt
	for _, v := range maxMap {
		if v < result {
			result = v
		}
	}
	if result == math.MaxInt {
		return 0
	}
	return result
}
