package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sumDistance([]int{-2, 0, 2}, "RLL", 3))
	// [-2,0,2], s = "RLL", d = 3
}
func sumDistance(nums []int, s string, d int) int {
	for i := 0; i < len(nums); i++ {
		if s[i] == 'L' {
			nums[i] -= d
		} else {
			nums[i] += d
		}
	}
	sort.Ints(nums)
	mod := int(1e9 + 7)
	var result int
	var sum = nums[0]
	for i := 1; i < len(nums); i++ {
		result += (nums[i]*i - sum) % mod
		result %= mod
		sum += nums[i]
		sum %= mod
	}
	result += mod
	result %= mod

	return result
}
