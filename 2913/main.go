package main

func main() {
}

func sumCounts(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		cntMap := make(map[int]int)
		for j := i; j < len(nums); j++ {
			cntMap[nums[j]]++
			result += len(cntMap) * len(cntMap)
		}
		result %= 1000000007
	}
	return result
}
