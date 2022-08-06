package main

import "fmt"

func main() {
	// fmt.Println(minimumReplacement([]int{19, 7, 2, 24, 11, 16, 1, 11, 23}))
	fmt.Println(minimumReplacement([]int{12, 9, 7, 6, 17, 19, 21}))
}

func minimumReplacement(nums []int) int64 {
	var result int64
	pre := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= pre {
			pre = nums[i]
			continue
		}
		cnt := nums[i] / pre
		if nums[i]%pre != 0 {
			cnt += 1
		}
		pre = nums[i] / cnt
		result += int64(cnt - 1)
	}
	return result
}
