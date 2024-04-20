package main

import "fmt"

func main() {
	fmt.Println(countAlternatingSubarrays([]int{0, 1, 1, 1}))
	fmt.Println(countAlternatingSubarrays([]int{1, 0, 1, 0}))
}

func countAlternatingSubarrays(nums []int) int64 {
	result := int64(0)
	pre := 0
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			result += int64((i - pre) * (i - pre + 1) / 2)
			pre = i
		}
	}
	result += int64((len(nums) - pre) * (len(nums) - pre + 1) / 2)
	return result
}
