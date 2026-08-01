package main

func main() {
}

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}
	if minVal == maxVal {
		return 0
	}

	bucketSize := (maxVal - minVal) / (n - 1)
	if bucketSize == 0 {
		bucketSize = 1
	}
	bucketNum := (maxVal-minVal)/bucketSize + 1

	bucketMin := make([]int, bucketNum)
	bucketMax := make([]int, bucketNum)
	hasValue := make([]bool, bucketNum)
	for i := 0; i < bucketNum; i++ {
		bucketMin[i] = 1<<31 - 1
		bucketMax[i] = -1 << 31
	}

	for _, num := range nums {
		idx := (num - minVal) / bucketSize
		if num < bucketMin[idx] {
			bucketMin[idx] = num
		}
		if num > bucketMax[idx] {
			bucketMax[idx] = num
		}
		hasValue[idx] = true
	}

	result := 0
	prevMax := minVal
	for i := 0; i < bucketNum; i++ {
		if !hasValue[i] {
			continue
		}
		if bucketMin[i]-prevMax > result {
			result = bucketMin[i] - prevMax
		}
		prevMax = bucketMax[i]
	}
	return result
}
