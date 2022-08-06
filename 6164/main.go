package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumSum([]int{18, 43, 36, 13, 7}))
}

func maximumSum(nums []int) int {
	maxMap := make(map[int][]int)
	for _, v := range nums {
		sum := getSumOfDigit(v)
		vals := maxMap[sum]
		vals = append(vals, v)
		if len(vals) <= 2 {
			maxMap[sum] = vals
			continue
		}
		sort.Ints(vals)
		maxMap[sum] = vals[len(vals)-2 : len(vals)]
	}
	result := -1
	for _, v := range maxMap {
		if len(v) < 2 {
			continue
		}
		val := v[0] + v[1]
		if val > result {
			result = val
		}
	}
	return result
}

func getSumOfDigit(num int) int {
	val := num
	res := 0
	for val > 0 {
		res += val % 10
		val /= 10
	}
	return res
}
