package main

import "fmt"

func main() {
	fmt.Println(isArraySpecial([]int{1}))
}

// The first element (index 0) must have an odd number of bits set in its binary representation.
// For each subsequent element, the number of bits set in its binary representation should alternate with the previous element.
func isArraySpecial(nums []int) bool {
	// Initialize a variable to hold the parity of the first element. The parity of a number is determined by its odd or even value.
	pre := nums[0] % 2

	// Iterate through each element in the array, starting from index 1 (we skip the first element as it has already been processed).
	for i, num := range nums {
		// If this is the first iteration, skip to the next statement.
		if i == 0 {
			continue
		}

		// Determine the parity of the current number.
		cur := num % 2

		// Check if the combined parity of the current and previous elements is not equal to 1 (indicating an inconsistency).
		if cur+pre != 1 {
			return false
		}

		// Update the variable that holds the parity of the previous element.
		pre = cur
	}

	// If we've successfully checked all elements without encountering any inconsistencies, return true (the array is special).
	return true
}
