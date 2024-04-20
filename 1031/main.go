package main

func main() {

}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	return max(helper(nums, firstLen, secondLen), helper(nums, secondLen, firstLen))
}

func helper(nums []int, firstLen int, secondLen int) int {
	maxs := make([]int, len(nums))
	firstCnt := 0
	var max int
	var sum int

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		firstCnt++
		if firstCnt < firstLen {
			maxs[i] = -1
			continue
		}
		if firstCnt > firstLen {
			sum -= nums[i-firstLen]
		}
		if sum > max {
			max = sum
		}
		maxs[i] = max
	}

	secondCnt := 0
	max = 0
	sum = 0
	var result int

	for i := len(nums) - 1; i >= 0; i-- {
		sum += nums[i]
		secondCnt++
		if secondCnt < secondLen {
			continue
		}
		if secondCnt > secondLen {
			sum -= nums[i+secondLen]
		}
		if sum > max {
			max = sum
		}
		if i >= firstLen && max+maxs[i-1] > result {
			result = sum + maxs[i-1]
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
