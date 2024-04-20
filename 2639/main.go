package main

import "strconv"

func main() {

}

func findColumnWidth(grid [][]int) []int {
	if len(grid) == 0 {
		return []int{}
	}

	var result []int

	for i := 0; i < len(grid[0]); i++ {
		var maxLenth int
		for j := 0; j < len(grid); j++ {
			str := strconv.Itoa(grid[j][i])
			if len(str) > maxLenth {
				maxLenth = len(str)
			}
		}
		result = append(result, maxLenth)
	}

	return result
}
