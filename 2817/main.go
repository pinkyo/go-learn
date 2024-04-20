package main

import (
	"sort"
)

func main() {
}

func minAbsoluteDifference(nums []int, x int) int {
	if x == 0 {
		return 0
	}
	numIdxArr := make([][]int, len(nums))
	for i, num := range nums {
		numIdxArr[i] = []int{num, i}
	}
	sort.Slice(numIdxArr, func(i, j int) bool { return numIdxArr[i][0] < numIdxArr[j][0] })
	minAbs := 1<<63 - 1
	for i := 0; i < len(numIdxArr)-1; i++ {
		j := i + 1
		for j < len(numIdxArr) && abs(numIdxArr[i][1]-numIdxArr[j][1]) < x && abs(numIdxArr[i][0]-numIdxArr[j][0]) < minAbs {
			j++
		}
		if j < len(numIdxArr) && abs(numIdxArr[i][0]-numIdxArr[j][0]) < minAbs {
			minAbs = abs(numIdxArr[i][0] - numIdxArr[j][0])
		}
		if minAbs == 0 {
			return 0
		}
	}
	return minAbs
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
