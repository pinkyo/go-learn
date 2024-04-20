package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(splitNum(687))
}

func splitNum(num int) int {
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Ints(digits)
	var inc int
	var result int
	exp := 1
	for i := len(digits) - 1; i >= 0; i -= 2 {
		val := digits[i] + inc
		inc = 0
		if i-1 >= 0 {
			val += digits[i-1]
		}
		if val > 10 {
			inc = 1
			val %= 10
		}
		result += val * exp
		exp *= 10
	}
	if inc > 0 {
		result += inc * exp
	}

	return result
}
