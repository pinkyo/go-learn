package main

func main() {
}

func maxFrequencyElements(nums []int) int {
	cntMap := make(map[int]int)
	for _, nums := range nums {
		cntMap[nums]++
	}

	maxVal := 0
	result := 0
	for _, cnt := range cntMap {
		if cnt > maxVal {
			result = cnt
			maxVal = cnt
		} else if cnt == maxVal {
			result += cnt
		}
	}

	return result
}
