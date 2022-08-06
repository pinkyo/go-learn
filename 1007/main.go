package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2}))
}

func minDominoRotations(tops []int, bottoms []int) int {
	cntMap := make(map[int]int)
	for i := 0; i < len(tops); i++ {
		if tops[i] == bottoms[i] {
			cntMap[tops[i]]++
			continue
		}
		cntMap[tops[i]]++
		cntMap[bottoms[i]]++
	}
	result := math.MaxInt
	topCntMap := make(map[int]int)
	bottomCntMap := make(map[int]int)
	for i := 0; i < len(tops); i++ {
		topCntMap[tops[i]]++
		bottomCntMap[bottoms[i]]++
	}
	for k, cnt := range cntMap {
		if cnt >= len(tops) {
			result = min(result, len(tops)-topCntMap[k])
			result = min(result, len(tops)-bottomCntMap[k])
		}
	}
	if result > len(tops) {
		return -1
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
