package main

import "sort"

func main() {

}

func maxStrength(nums []int) int64 {
	sort.Ints(nums)
	max := int64(1)
	var cnt int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			max *= int64(nums[i])
			cnt++
		}
		if nums[i] == 0 {
			continue
		}
		if nums[i] < 0 && i+1 < len(nums) && nums[i+1] < 0 {
			max *= int64(nums[i] * nums[i+1])
			cnt += 2
			i++
		}
	}
	if cnt == 0 && len(nums) == 0 {
		return 0
	} else if cnt == 0 {
		return int64(nums[len(nums)-1])
	}
	return max
}
