package main

import (
	"fmt"
)

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 1}, {1, 0},
	}))
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 || grid[len(grid)-1][len(grid[0])-1] != 0 {
		return -1
	}

	list := make([][]int, 0)
	list = append(list, []int{0, 0, 1})
	grid[0][0] = 1

	addItem := func(i, j, depth int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) || grid[i][j] != 0 {
			return
		}
		grid[i][j] = 1
		list = append(list, []int{i, j, depth})
	}
	for len(list) > 0 {
		head := list[0]
		if head[0] == len(grid)-1 && head[1] == len(grid[0])-1 {
			return head[2]
		}
		addItem(head[0]+1, head[1]+1, head[2]+1)
		addItem(head[0], head[1]+1, head[2]+1)
		addItem(head[0]+1, head[1], head[2]+1)
		addItem(head[0]-1, head[1]+1, head[2]+1)
		addItem(head[0]+1, head[1]-1, head[2]+1)
		addItem(head[0]-1, head[1], head[2]+1)
		addItem(head[0], head[1]-1, head[2]+1)
		addItem(head[0]-1, head[1]-1, head[2]+1)
		list = list[1:]
	}
	return -1
}
