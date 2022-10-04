package main

func main() {
}

func maxSubarraySumCircular(nums []int) int {
	left := make([]int, len(nums))
	max := nums[0]
	sum := nums[0]
	left[0] = nums[0]
	cur := left[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		left[i] = left[i-1]
		if sum > left[i] {
			left[i] = sum
		}
		cur += nums[i]
		if cur < nums[i] {
			cur = nums[i]
		}
		if cur > max {
			max = cur
		}
	}
	rightMax := nums[len(nums)-1]
	sum2 := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if rightMax+left[i] > max {
			max = rightMax + left[i]
		}
		sum2 += nums[i]
		if sum2 > rightMax {
			rightMax = sum2
		}
	}

	return max
}
