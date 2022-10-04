package main

import "fmt"

func main() {
	fmt.Println(smallestSubarrays([]int{1, 0, 2, 1, 3}))
}

func smallestSubarrays(nums []int) []int {
	pos := make([]int, 64)
	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		val := nums[i]
		var idx int
		for val > 0 {
			if val&1 == 1 {
				pos[idx] = i
			}
			val >>= 1
			idx++
		}
		res := 1
		for _, p := range pos {
			if p-i+1 > res {
				res = p - i + 1
			}
		}
		result[i] = res
	}
	return result
}
