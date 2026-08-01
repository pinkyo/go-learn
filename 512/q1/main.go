package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestInteger(1, 1))
}

func largestInteger(n int, s int) int {
	if s > n*9 {
		return -1
	}
	result := 0
	for i := 0; i < n; i++ {
		if s >= 9 {
			result += 9 * int(math.Pow(10, float64(n-i-1)))
			s -= 9
		} else if s > 0 {
			result += s * int(math.Pow(10, float64(n-i-1)))
			s = 0
		}
	}
	return result
}
