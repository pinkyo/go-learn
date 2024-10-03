package main

import "fmt"

func main() {
	//	fmt.Println(checkXMatrix([][]int{
	//		{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2},
	//	}))
	//
	printlnA(1)
}

// Function to check if a matrix is X-Matrix
func checkXMatrix(grid [][]int) bool {
	// Iterate through each row of the matrix
	for i := 0; i < len(grid); i++ {
		// Iterate through each column of the matrix
		for j := 0; j < len(grid[0]); j++ {
			// If the row and column indices are equal or the row index is equal to the length of the row minus the column index minus 1
			if i == j || i == len(grid[i])-j-1 {
				// If the value at this index is 0, return false
				if grid[i][j] == 0 {
					return false
				}
				// Otherwise, if the value at this index is not 0, return false
			} else if grid[i][j] != 0 {
				return false
			}
		}
	}
	// If all conditions are met, return true
	return true
}

// This function prints the value of the integer 'a'
func printlnA(a int) {
	// Print the value of 'a'
	fmt.Println(a)
}
