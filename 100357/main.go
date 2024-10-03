package main

import "fmt"

// This function is the main function of the program
func main() {
	// This line prints the maximum length of the array [1, 2, 3, 4]
	fmt.Println(maximumLength([]int{1, 2, 3, 4}))
	// This line prints the maximum length of the array [1, 2, 1, 1, 2, 1, 2]
	fmt.Println(maximumLength([]int{1, 2, 1, 1, 2, 1, 2}))
}

// This function takes a slice of integers as an argument and returns the maximum length of consecutive numbers of the same parity.
func maximumLength(nums []int) int {
	// Initialize the result to 0
	result := 0
	// Initialize a slice of integers to store the counts of 0s and 1s
	values := [2]int{0, 0}
	// Iterate through the slice of integers
	for _, num := range nums {
		// Increment the count of 0s or 1s based on the remainder of the integer divided by 2
		values[num%2]++
	}
	// Update the result to the maximum of the counts of 0s and 1s
	result = max(values[0], values[1])
	// Initialize the slice of integers to store the counts of 0s and 1s
	values = [2]int{0, 0}
	// Iterate through the slice of integers
	for _, num := range nums {
		// If the remainder of the integer divided by 2 matches the parity of the current count of 0s, increment the count of 0s
		if num%2 == values[0]%2 {
			values[0]++
		}
		// If the remainder of the integer divided by 2 does not match the parity of the current count of 1s, increment the count of 1s
		if num%2 != values[1]%2 {
			values[1]++
		}
	}
	// Update the result to the maximum of the counts of 0s and 1s
	result = max(result, values[0])
	result = max(result, values[1])

	// Return the result
	return result
}

// This function takes two integers as parameters and returns the larger of the two.
func max(a, b int) int {
	// If the first integer is greater than the second, return the first integer.
	if a > b {
		return a
	}
	// Otherwise, return the second integer.
	return b
}
