package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	fmt.Println(beautifulIndices(
		"sqgrt",
		"rt",
		"sq",
		3,
	))
}

func beautifulIndices(s string, a string, b string, k int) []int {
	preB := -1
	ans := []int{}
	for i := 0; i < len(s); i++ {
		bMatch := strings.HasPrefix(s[i:], b)
		if bMatch {
			preB = i
		}
		aMatch := strings.HasPrefix(s[i:], a)
		if aMatch && preB >= 0 && abs(i-preB) <= k {
			ans = append(ans, i)
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		bMatch := strings.HasPrefix(s[i:], b)
		if bMatch {
			preB = i
		}
		aMatch := strings.HasPrefix(s[i:], a)
		if aMatch && preB >= 0 && abs(i-preB) <= k {
			ans = append(ans, i)
		}
	}

	if len(ans) == 0 {
		return []int{}
	}

	// Deduplicate and sort
	sort.Ints(ans)
	result := []int{}
	result = append(result, ans[0])
	for i := 1; i < len(ans); i++ {
		if ans[i] != ans[i-1] {
			result = append(result, ans[i])
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a ...int) int {
	min := math.MaxInt32
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

// Returns the maximum value from the given list of integers.
func max(a ...int) int {
	max := math.MinInt32
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
