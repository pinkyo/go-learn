package main

import "fmt"

func main() {
	// fmt.Println(maxMoves([][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}))
	fmt.Println(maxMoves([][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}))
}

func maxMoves(grid [][]int) int {
	array := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		array[i] = make([]int, len(grid[0]))
	}

	var result int
	for i := 0; i < len(grid[0])-1; i++ {
		for j := 0; j < len(grid); j++ {
			if i != 0 && array[j][i] <= 0 {
				continue
			}
			if j > 0 && grid[j-1][i+1] > grid[j][i] {
				array[j-1][i+1] = max(array[j-1][i+1], array[j][i]+1)
				result = max(result, array[j-1][i+1])
			}
			if grid[j][i+1] > grid[j][i] {
				array[j][i+1] = max(array[j][i+1], array[j][i]+1)
				result = max(result, array[j][i+1])
			}
			if j < len(grid)-1 && grid[j+1][i+1] > grid[j][i] {
				array[j+1][i+1] = max(array[j+1][i+1], array[j][i]+1)
				result = max(result, array[j+1][i+1])
			}
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
