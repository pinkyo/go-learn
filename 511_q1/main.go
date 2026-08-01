package main

import "fmt"

func main() {
	fmt.Println(canReach([]int{1, 1}, []int{2, 2}))
	fmt.Println(canReach([]int{4, 5}, []int{6, 6}))
}

func canReach(start []int, target []int) bool {
	chessboard := [8][8]int{}

	chessboard[start[0]][start[1]] = 1
	arr := [][]int{
		{-1, -2},
		{-2, -1},
		{-2, 1},
		{-1, 2},
		{1, 2},
		{2, 1},
		{2, -1},
		{1, -2},
	}
	arr2 := [][]int{}
	for _, v := range arr {
		for _, v2 := range arr {
			arr2 = append(arr2, []int{v[0] + v2[0], v[1] + v2[1]})
		}
	}

	queue := [][]int{start}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, dir := range arr2 {
			next := []int{cur[0] + dir[0], cur[1] + dir[1]}
			if next[0] < 0 || next[0] >= 8 || next[1] < 0 || next[1] >= 8 {
				continue
			}
			if chessboard[next[0]][next[1]] == 1 {
				continue
			}
			chessboard[next[0]][next[1]] = 1
			queue = append(queue, next)
		}
	}

	return chessboard[target[0]][target[1]] == 1
}
