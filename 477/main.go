package main

import "sort"

func main() {
}

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = num * num
	}

	sort.Ints(result)
	return result
}
