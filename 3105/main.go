package main

func main() {
}

func longestMonotonicSubarray(nums []int) int {
	result := 0
	increase, decrease := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			result = max(result, decrease)
			decrease = 1
			increase += 1
		} else if nums[i] < nums[i-1] {
			result = max(result, increase)
			increase = 1
			decrease += 1
		} else {
			result = max(result, increase, decrease)
			increase, decrease = 1, 1
		}
	}
	result = max(result, increase, decrease)
	return result
}

func max(any ...int) int {
	result := any[0]
	for _, v := range any {
		if result < v {
			result = v
		}
	}
	return result
}
