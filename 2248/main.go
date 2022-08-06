package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(intersection(
		[][]int{
			{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6},
		},
	))
}

func intersection(nums [][]int) []int {
	cntMap := make(map[int]int)
	for _, v := range nums {
		for _, vv := range v {
			cntMap[vv]++

		}
	}
	var result []int
	for k, v := range cntMap {
		if v == len(nums) {
			result = append(result, k)
		}
	}
	sort.Ints(result)
	return result
}
