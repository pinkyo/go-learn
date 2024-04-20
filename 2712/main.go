package main

import "fmt"

func main() {
	fmt.Println(minimumCost("0011"))
	fmt.Println(minimumCost("010101"))
	fmt.Println(minimumCost("000000010"))
}

func minimumCost(s string) int64 {
	var result0 int
	var result1 int

	var cnt1 int
	var cnt0 int
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			cnt0++
		} else {
			cnt1++
		}
		if i > 0 && s[i] != s[i-1] {
			if s[i] == '1' {
				result0 += min(i*2-int(cnt0), (len(s)-i)*2+cnt0)
				cnt0 = 0
			} else {
				result1 += min(i*2-int(cnt1), (len(s)-i)*2+cnt1)
				cnt1 = 0
			}
		}
	}
	if cnt1 > 0 {
		result1 += cnt1
	}
	if cnt0 > 0 {
		result0 += cnt0
	}
	return int64(min(result0, result1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
