package main

import "sort"

func main() {
}

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	for i := 0; i < len(moveFrom); i++ {
		if moveFrom[i] == moveTo[i] {
			continue
		}
		numMap[moveTo[i]] += numMap[moveFrom[i]]
		numMap[moveFrom[i]] = 0
	}
	result := make([]int, 0, len(nums))
	for k, v := range numMap {
		if v > 0 {
			result = append(result, k)
		}
	}
	sort.Ints(result)
	return result
}
