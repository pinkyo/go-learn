package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(rearrangeCharacters("ilovecodingonleetcode", "code"))
}

func rearrangeCharacters(s string, target string) int {
	cntArr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		cntArr[s[i]-'a']++
	}
	cntArr2 := make([]int, 26)
	for i := 0; i < len(target); i++ {
		cntArr2[target[i]-'a']++
	}
	result := math.MaxInt64
	for i := 0; i < len(cntArr2); i++ {
		if cntArr2[i] == 0 {
			continue
		}
		result = min(result, cntArr[i]/cntArr2[i])
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
