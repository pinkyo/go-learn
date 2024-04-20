package main

import "fmt"

func main() {
	fmt.Println(maxSum([]int{51, 71, 17, 24, 42}))
}

func maxSum(nums []int) int {
	result := -1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if maxOfDigit(nums[i]) == maxOfDigit(nums[j]) && result < nums[i]+nums[j] {
				result = nums[i] + nums[j]
			}
		}
	}

	return result
}

func maxOfDigit(val int) int {
	result := 0
	for val > 0 {
		if result < val%10 {
			result = val % 10
		}
		val /= 10
	}
	return result
}
