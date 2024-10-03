package main

import "fmt"

func main() {
	fmt.Println(longestSubsequence([]int{3, 4, -3, -2, -4}, -5))
}

func longestSubsequence(arr []int, difference int) int {
	dp := make([]int, len(arr))
	preMap := make(map[int]int, 0)

	for i := 0; i < len(arr); i++ {
		val := arr[i]
		pre := arr[i] - difference
		if preIdx, ok := preMap[pre]; ok {
			dp[i] = dp[preIdx] + 1
		} else {
			dp[i] = 1
		}
		preMap[val] = i
	}

	result := 0
	for _, val := range dp {
		if val > result {
			result = val
		}
	}

	return result
}
