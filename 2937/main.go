package main

import "fmt"

func main() {
	fmt.Println(findMinimumOperations("aca", "abcc", "accba"))
}

func findMinimumOperations(s1 string, s2 string, s3 string) int {
	result := 0
	str := calMaxCommonStringHelper(s1, s2)
	if str == "" {
		return -1
	}
	result += len(s1) + len(s2) - 2*len(str)
	str2 := calMaxCommonStringHelper(str, s3)
	if str2 == "" {
		return -1
	}
	result += len(s3) + 2*len(str) - 3*len(str2)
	return result
}

func calMaxCommonStringHelper(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 || s1[0] != s2[0] {
		return ""
	}
	dp := make([][]int, len(s1)+1)
	str := make([][]string, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
		str[i] = make([]string, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			dp[i][j] = dp[i-1][j]
			str[i][j] = str[i-1][j]
			if dp[i][j-1] > dp[i][j] {
				dp[i][j] = dp[i][j-1]
				str[i][j] = str[i][j-1]
			}
			if s1[i-1] == s2[j-1] && dp[i-1][j-1]+1 > dp[i][j] {
				dp[i][j] = dp[i-1][j-1] + 1
				str[i][j] = str[i-1][j-1] + string(s1[i-1])
			}
		}
	}

	// 找到最长公共子串
	return str[len(s1)][len(s2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
