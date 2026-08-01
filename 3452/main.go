package main

func main() {
}

func sumOfGoodNumbers(nums []int, k int) int {
	result := 0
	for i, num := range nums {
		if (i < k || nums[i] > nums[i-k]) && (i+k >= len(nums) || nums[i] > nums[i+k]) {
			result += num
		}
	}

	return result
}
