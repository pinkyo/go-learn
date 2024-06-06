package main

import "fmt"

func main() {
	fmt.Println(maxScore([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}

func maxScore(grid [][]int) int {
	// Initialize the result as the difference between the first element and the second element in the first row
	result := grid[0][1] - grid[0][0]
	// Iterate over the entire grid using nested loops
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// Skip the starting point
			if i == 0 && j == 0 {
				continue
			}
			// Initialize the minimum value at the current position as positive infinity
			minVal := int(1e9)
			// If not on the first row, update the minimum value with the left neighbor's value
			if i > 0 {
				minVal = grid[i-1][j]
			}
			// If not on the first column, update the minimum value with the upper neighbor's value
			if j > 0 {
				minVal = min(minVal, grid[i][j-1])
			}
			// If the current position's value minus the minimum value is greater than the current highest score, update the highest score
			if grid[i][j]-minVal > result {
				result = grid[i][j] - minVal
			}
			// If the current position's value is greater than the minimum value, update the current position's value to the minimum value
			if grid[i][j] > minVal {
				grid[i][j] = minVal
			}
		}
	}
	return result
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
