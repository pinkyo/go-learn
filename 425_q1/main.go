package main

func main() {
}

func minimumSumSubarray(nums []int, l int, r int) int {
	result := -1
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j	++ {
			sum += nums[j]
			if j-i+1 < l {
				continue
			}
			if j-i+1 > r {
				break
			}
			if sum > 0 && (result == -1 || sum < result) {
				result = sum
			}
		}
	}

	return result
}
