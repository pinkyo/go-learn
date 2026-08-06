package main

import "math"

func main() {
}

func countValidPrefixes(s string) int {
	cnt0, cnt1 := 0, 0
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			cnt0++
		} else {
			cnt1++
		}
		if math.Abs(float64(cnt0-cnt1)) <= 1 {
			result++
		}
	}
	return result
}
