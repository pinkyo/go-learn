package main

func main() {

}

func differenceOfDistinctValues(grid [][]int) [][]int {
	result := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		result[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			topLeft := make(map[int]struct{})
			bottomRight := make(map[int]struct{})
			for k := 1; k <= i && k <= j; k++ {
				topLeft[grid[i-k][j-k]] = struct{}{}
			}
			for k := 1; k < len(grid)-i && k < len(grid[i])-j; k++ {
				bottomRight[grid[i+k][j+k]] = struct{}{}
			}
			result[i][j] = abs(len(topLeft) - len(bottomRight))
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
