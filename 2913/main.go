package main

import "fmt"

func main() {
	// Test case
	fmt.Println(sumCounts([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

// This function calculates the sum of the squares of the distinct counts of elements in all subarrays starting from each index.
//
// Parameters:
//   - nums: A slice of integers for which the sum of squares of distinct counts of elements in subarrays is calculated
//
// Returns:
//   - int: The result of the calculation modulo 1000000007
func sumCounts(nums []int) int {
	var result int // Initialize the result variable to store the final sum

	// Iterate over each starting index of the subarray
	for i := 0; i < len(nums); i++ {
		cntMap := make(map[int]int) // Create a map to count occurrences of each number in the subarray

		// Iterate over each ending index of the subarray starting from index i
		for j := i; j < len(nums); j++ {
			cntMap[nums[j]]++                   // Increment the count of the current number in the map
			result += len(cntMap) * len(cntMap) // Add the square of the number of unique elements to the result
		}
		result %= 1000000007 // Take modulo 1000000007 to prevent integer overflow
	}
	return result // Return the final result
}
