package main

import "fmt"

func main() {
	fmt.Println(longestSemiRepetitiveSubstring("52233"))
	fmt.Println(longestSemiRepetitiveSubstring("5494"))
	fmt.Println(longestSemiRepetitiveSubstring("1111111"))
	fmt.Println(longestSemiRepetitiveSubstring("011"))
	fmt.Println(longestSemiRepetitiveSubstring("0001"))
}

func longestSemiRepetitiveSubstring(s string) int {
	if len(s) <= 2 {
		return len(s)
	}
	result := 2
	preArr := [2]int{-1, -1}
	for i := 1; i < len(s); i++ {
		val := s[i]
		if val == s[i-1] {
			if preArr[1] >= 0 {
				result = max(result, i-preArr[1])
			} else {
				result = max(result, i+1)
			}
			preArr[0] = preArr[1]
			preArr[1] = i
			continue
		}
		if preArr[0] == -1 {
			result = max(result, i+1)
		} else {
			result = max(result, i-preArr[0]+1)
		}

	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
