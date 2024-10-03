package main

import "fmt"

func main() {
	fmt.Println(maximumPoints([]int{2}, 10))
}

// Function to find the maximum points
//
// Parameters:
//   - enemyEnergies: list of enemy energies
//   - currentEnergy: current energy of the players
//
// Returns:
//   - maximum points
func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
	// Initialize minVal to a very large number
	minVal := int(1e9 + 1)
	// Initialize sum with currentEnergy
	sum := int(currentEnergy)
	// Iterate through each enemy's energy
	for _, enemy := range enemyEnergies {
		// Update minVal if the current enemy's energy is smaller
		minVal = min(minVal, enemy)
		// Add the current enemy's energy to sum
		sum += enemy
	}
	// If currentEnergy is less than minVal, return 0
	if currentEnergy < int(minVal) {
		return 0
	}
	// Calculate and return the maximum points
	return int64((sum - minVal) / minVal)
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
