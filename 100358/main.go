package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumLength([]int{1, 4, 2, 3, 1, 4}, 3))
}

func maximumLength(nums []int, k int) int {
	result := 0

	// for r := 0; r < k; r++ {
	// 	values := make([]int, k)
	// 	from := make([]int, k)
	// 	for c := 0; c < k; c++ {
	// 		from[c] = c
	// 	}
	// 	for _, num := range nums {
	// 		for i := 0; i < k; i++ {
	// 			if from[i] == num%k {
	// 				values[i]++
	// 				from[i] = (r + k - from[i]) % k
	// 			}
	// 		}
	// 	}
	// 	for _, value := range values {
	// 		result = max(result, value)
	// 	}
	// }

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
