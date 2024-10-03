package main

import "fmt"

func main() {
	// fmt.Println(numberOfAlternatingGroups([]int{0, 0, 1}, 3))
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 1}, 3))
}

// This function counts the number of alternating groups in a given array of colors
// The function takes in two parameters: an array of colors and an integer k
func numberOfAlternatingGroups(colors []int, k int) int {
	// Initialize a counter to 1
	cnt := 1
	// Initialize a result variable to 0
	result := 0
	// Loop through the array of colors
	for i := 1; i < len(colors)+k-1; i++ {
		// If the current color is not equal to the previous color
		if colors[i%len(colors)] != colors[(i-1)%len(colors)] {
			// Increment the counter
			cnt++
		} else {
			// If the counter is greater than or equal to k
			if cnt >= k {
				// Add the difference between the counter and k to the result
				result += cnt - k + 1
			}
			// Reset the counter to 1
			cnt = 1
		}
	}
	// If the counter is greater than or equal to k
	if cnt >= k {
		// Add the difference between the counter and k to the result
		result += cnt - k + 1
	}
	// Return the result
	return result
}
