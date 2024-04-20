package main

func main() {
}

func maxArrayValue(nums []int) int64 {
	var result int64
	cur := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if cur == 0 {
			cur = nums[i]
		} else if nums[i] <= cur {
			cur += nums[i]
		} else if cur > int(result) {
			result = int64(cur)
			cur = nums[i]
		}
	}
	if cur > int(result) {
		result = int64(cur)
	}
	return result
}
