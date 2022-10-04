package main

import "fmt"

func main() {
	// fmt.Println(maximumXOR([]int{3, 2, 4, 6}))
	fmt.Println("Hello, playground")
}

func maximumXOR(nums []int) int {
	hasArr := make([]bool, 32)
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		idx := 0
		for val > 0 {
			if val&1 == 1 {
				hasArr[idx] = true
			}
			idx++
			val >>= 1
		}
	}
	var result int
	for i := 0; i < 32; i++ {
		if hasArr[i] {
			result |= 1 << i
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
