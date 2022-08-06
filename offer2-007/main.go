package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	resultMap := make(map[[3]int]struct{})
	cntMap := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			c := -(nums[i] + nums[j])
			if _, ok := cntMap[c]; ok {
				slice := []int{nums[i], nums[j], c}
				sort.Ints(slice)
				resultMap[[3]int{slice[0], slice[1], slice[2]}] = struct{}{}
			}
		}
		cntMap[nums[i]] = struct{}{}
	}
	var result [][]int
	for k := range resultMap {
		result = append(result, []int{k[0], k[1], k[2]})
	}
	return result
}
