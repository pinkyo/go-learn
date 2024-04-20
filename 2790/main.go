package main

import "fmt"

func main() {
	fmt.Println(maxIncreasingGroups([]int{2, 1, 2}))
}

func maxIncreasingGroups(usageLimits []int) int {
	sum := 0
	for _, l := range usageLimits {
		sum += l
	}
	result := len(usageLimits)

	if result*(result+1)/2 > sum {
		result = cal(sum)
	}
	return result
}

func cal(sum int) int {
	l, h := 0, sum
	for l <= h {
		m := (l + h) / 2
		if m*(m+1)/2 > sum {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return h
}
