package main

func main() {
}

func findKOr(nums []int, k int) int {
	var result int
	base := 1
	for {
		var count int
		var nonZeroCount int
		for i := 0; i < len(nums); i++ {
			if nums[i] == 0 {
				continue
			}
			nonZeroCount++
			if nums[i]%2 == 1 {
				count++
			}
			nums[i] /= 2
		}
		if nonZeroCount == 0 {
			break
		} else if count >= k {
			result += base
		}
		base <<= 1
	}
	return result
}
