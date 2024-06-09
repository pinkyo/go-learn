package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 1, 1}, 10))
}

// Given an integer array nums and a limit, return the length of the longest continuous subarray in nums such that the absolute difference between any two elements in this subarray is less than or equal to limit. If there's no subarray meeting this condition, return 0.

// Function to find the maximum value within a range
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Function to find the minimum value within a range
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Function longestSubarray implements a sliding window approach using two pointers (left and right) for the start and end of the current subarray, respectively
func longestSubarray(nums []int, limit int) int {
	// Initialize variables
	left, right := 0, 1 // left and right pointer for sliding window
	maxLength := 1      // max length of any valid subarray will be initialized to 1 (the length of the initial subarray)

	minValue := nums[left] // initialize minimum value in current subarray
	maxValue := nums[left] // initialize maximum value in current subarray
	// Loop through the array, incrementing the right pointer as we go
	for right < len(nums) {
		// Update maxValue and minValue with current value at right index
		maxValue = max(maxValue, nums[right])
		minValue = min(minValue, nums[right])
		// Check if the difference between maxValue and minValue is within limit
		newDiff := math.Abs(float64(maxValue - minValue))
		if newDiff <= float64(limit) {
			// If it is, increment maxLength and move left pointer forward
			maxLength = max(maxLength, right-left+1)
			right++
		} else {
			// If it's not, we need to find a new subarray that satisfies the condition. Move the right pointer forward until we do.
			left++
			if left >= len(nums) {
				break
			}
			for right-left+1 <= maxLength {
				right++
			}
			if right >= len(nums) {
				break
			}

			// Iterate over all elements from left to right, updating minValue and maxValue as needed
			for i := left; i <= right; i++ {
				minValue = min(minValue, nums[i])
				maxValue = max(maxValue, nums[i])
			}
		}
	}

	// Return the maximum length of any valid subarray found
	return maxLength
}
