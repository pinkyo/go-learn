package main

import "fmt"

func main() {
	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
}

func zeroFilledSubarray(nums []int) int64 {
	cnt := 0
	var result int64
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt++
			continue
		}
		result += int64(cnt * (cnt + 1) / 2)
		cnt = 0
	}
	result += int64(cnt * (cnt + 1) / 2)
	return result
}
