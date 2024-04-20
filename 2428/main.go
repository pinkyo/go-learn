package main

func main() {
}

func maxSum(grid [][]int) int {
	var result int
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			temp := grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if temp > result {
				result = temp
			}
		}
	}
	return result
}
