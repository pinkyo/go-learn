package main

func main() {
}

func countBadPairs(nums []int) int64 {
	cntMap := make(map[int]int)
	for i, num := range nums {
		cntMap[num-i]++
	}
	result := int64(len(nums) * (len(nums) - 1) / 2)
	for _, cnt := range cntMap {
		result -= int64((cnt - 1) * cnt / 2)
	}
	return result
}
