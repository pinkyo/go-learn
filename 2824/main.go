package main

import "fmt"

func main() {
}

func countPairs(nums []int, target int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			fmt.Println(len(nums))
			if nums[i]+nums[j] < target {
				result++
			}
		}
	}
	return result
}
