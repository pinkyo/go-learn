package main

import "fmt"

func main() {
//	fmt.Println(checkXMatrix([][]int{
//		{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2},
//	}))
    printlnA(1)
}

func checkXMatrix(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == j || i == len(grid[i])-j-1 {
				if grid[i][j] == 0 {
					return false
				}
			} else if grid[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func printlnA(a int) {
    fmt.Println(a)
}
