package main

import (
	"math"
	"strconv"
)

func main() {
}

func splitIntoFibonacci(num string) []int {
	var dfs func(num string, index int, fib []int) ([]int, bool)
	dfs = func(num string, index int, fib []int) ([]int, bool) {
		if index == len(num) {
			return fib, len(fib) >= 3
		}
		sum := fib[len(fib)-1] + fib[len(fib)-2]
		if sum > math.MaxInt32 {
			return fib, false
		}
		sumStr := strconv.Itoa(sum)
		if index+len(sumStr) > len(num) || num[index:index+len(sumStr)] != sumStr {
			return fib, false
		}
		return dfs(num, index+len(sumStr), append(fib, sum))
	}

	var first int
	for i := 0; i < len(num); i++ {
		first = first*10 + int(num[i]-'0')
		if first > math.MaxInt32 {
			break
		}
		var second int
		for j := i + 1; j < len(num); j++ {
			second = second*10 + int(num[j]-'0')
			if first > math.MaxInt32 {
				break
			}
			fib := []int{first, second}
			if res, ok := dfs(num, j+1, fib); ok {
				return res
			}
			if second == 0 {
				break
			}
		}
		if first == 0 {
			break
		}
	}
	return []int{}
}
