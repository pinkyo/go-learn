package main

import "fmt"

func main() {
	fmt.Println(findSubstringInWraproundString("zip"))
}

// Finding the longest substring that is a wrap around string.
func findSubstringInWraproundString(p string) int {
	cntArr := make([]int, 26)
	pre := 0
	cntArr[p[0]-'a'] = 1
	for i := 1; i < len(p); i++ {
		if p[i]-'a' == (p[i-1]-'a'+1)%26 {
			continue
		}
		for j := 0; j < 26 && j < i-pre; j++ {
			idx := (j + int(p[pre]-'a')) % 26
			cntArr[idx] = max(cntArr[idx], i-pre-j)
		}
		pre = i
	}
	for j := 0; j < 26 && j < len(p)-pre; j++ {
		idx := (j + int(p[pre]-'a')) % 26
		cntArr[idx] = max(cntArr[idx], len(p)-pre-j)
	}
	var result int
	for i := 0; i < len(cntArr); i++ {
		result += cntArr[i]
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
