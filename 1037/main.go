package main

import "fmt"

func main() {
	fmt.Println(isBoomerang([][]int{
		{1, 1}, {2, 3}, {3, 2},
	}))
}

func isBoomerang(points [][]int) bool {
	if isSame(points[0], points[1]) || isSame(points[1], points[2]) || isSame(points[0], points[2]) {
		return false
	}
	return (points[0][0]-points[1][0])*(points[1][1]-points[2][1]) !=
		(points[0][1]-points[1][1])*(points[1][0]-points[2][0])
}

func isSame(p1, p2 []int) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}
