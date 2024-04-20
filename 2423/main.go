package main

import "fmt"

func main() {
	// fmt.Println(equalFrequency("abcc"))
	// fmt.Println(equalFrequency("abbcc"))
	fmt.Println(equalFrequency("zz"))
}

func equalFrequency(word string) bool {
	cntArr := make([]int, 26)
	for _, c := range word {
		cntArr[c-'a']++
	}
	cntMap := make(map[int]int)
	for _, cnt := range cntArr {
		if cnt == 0 {
			continue
		}
		cntMap[cnt]++
	}
	if _, ok := cntMap[1]; ok && len(cntMap) == 1 {
		return true
	}
	var keys []int
	for k := range cntMap {
		keys = append(keys, k)
	}
	if len(keys) == 1 && cntMap[keys[0]] == 1 {
		return true
	}
	if len(cntMap) > 2 {
		return false
	}
	var val1, val2 int
	var cnt1, cnt2 int
	for val, cnt := range cntMap {
		if cnt1 == 0 {
			val1 = val
			cnt1 = cnt
		} else {
			val2 = val
			cnt2 = cnt
		}
	}
	return cnt1 == 1 && (val1 == 1 || val1 == val2+1) || cnt2 == 1 && (val2 == 1 || val2 == val1+1)
}
