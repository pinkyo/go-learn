package main

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1:
	fmt.Println(minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
	// Test case 2:
	fmt.Println(minimizeMax([]int{4, 2, 1, 2}, 1))
}

// This function takes a slice of integers and an integer p as input
// and returns an integer as output. It is used to minimize the maximum
// difference between two adjacent elements in the slice by removing at
// most p elements.
func minimizeMax(nums []int, p int) int {
	// If p is 0, return 0
	if p == 0 {
		return 0
	}
	// Sort the slice of integers in ascending order
	sort.Ints(nums)
	// Initialize l and h to 0 and the maximum difference between the
	// smallest and largest elements in the slice, respectively
	l, h := 0, nums[len(nums)-1]-nums[0]
	// Loop until l is less than or equal to h
	for l <= h {
		// Calculate the mid value
		mid := (l + h) / 2
		// If check(nums, mid, p) is true, set h to mid - 1
		if check(nums, mid, p) {
			h = mid - 1
			// Otherwise, set l to mid + 1
		} else {
			l = mid + 1
		}
	}
	// Return l
	return l
}

// This function takes a slice of integers, a mid value, and an integer p
// as input and returns a boolean value as output. It is used to check if
// the maximum difference between two adjacent elements in the slice can
// be less than or equal to mid by removing at most p elements.
func check(nums []int, mid, p int) bool {
	// Initialize cnt to 0
	var cnt int
	// Loop through the slice of integers starting from the second element
	for i := 1; i < len(nums); i++ {
		// If the difference between the current element and the previous
		// element is less than or equal to mid, increment cnt and skip the
		// next element
		if nums[i]-nums[i-1] <= mid {
			cnt++
			i++
		}
		// If cnt is greater than or equal to p, return true
		if cnt >= p {
			return true
		}
	}
	// Return false
	return false
}
