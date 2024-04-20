package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([]int{2, 1, 3, 2, 1}))
	fmt.Println(minimumOperations([]int{2, 2, 2, 2, 3, 3}))
}

func minimumOperations(nums []int) int {
	return -1
	// max1 := 0
	// min3
	// for i := 0; i < len(nums); i++ {
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
