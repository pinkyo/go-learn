package main

import "fmt"

func main() {
	fmt.Println(maximumEnergy([]int{-2, -3, -1}, 2)) // Output: -1
}

// Function to calculate the maximum energy that can be gained during a move (k steps) in an array of energies.
func maximumEnergy(energy []int, k int) int {
	// Iterate over the energy array from the kth index to the end, adding the energy at the previous k indices if it's positive.
	for i := k; i < len(energy); i++ {
		if energy[i-k] > 0 {
			energy[i] += energy[i-k] // Add the previous energy to the current energy.
		}
	}

	// Initialize the result variable with the last element of the modified energy array (which represents the maximum energy that can be gained during one move).
	result := energy[len(energy)-1]

	// Iterate over the energy array from the second last index to the first index before k, comparing the current result with the energies at these indices. If the current result is smaller, update it.
	for i := 1; i < k; i++ {
		if result < energy[len(energy)-i-1] { // Check if the result is smaller than the energy at the previous index.
			result = energy[len(energy)-i-1] // If so, update the result with the energy at the previous index.
		}
	}

	// Return the final result (the maximum energy that can be gained during a move).
	return result
}
