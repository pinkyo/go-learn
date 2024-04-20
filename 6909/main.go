package main

func main() {
}

func longestAlternatingSubarray(nums []int, threshold int) int {
	var result int
	var cnt int
	for i := range nums {
		if i > 0 && nums[i]%2 != nums[i-1]%2 && nums[i] <= threshold && cnt > 0 {
			cnt++
			continue
		}
		if cnt > result {
			result = cnt
		}
		cnt = 0
		if nums[i]%2 == 0 && nums[i] <= threshold && cnt == 0 {
			cnt++
		}
	}
	if cnt > result {
		result = cnt
	}
	return result
}
