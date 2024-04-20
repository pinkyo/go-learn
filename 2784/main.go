package main

import "sort"

func main() {
}

func isGood(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != i+1 {
			return false
		}
	}
	return nums[len(nums)-1] == nums[len(nums)-2]
}
