package main

import "sort"

func main() {
}

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)
	midIdx := len(nums) / 2
	result := int64(abs(nums[midIdx] - k))
	for i := midIdx - 1; i >= 0 && nums[i] > k; i-- {
		result += int64(nums[i] - k)
	}
	for i := midIdx + 1; i < len(nums) && nums[i] < k; i++ {
		result += int64(k - nums[i])
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
