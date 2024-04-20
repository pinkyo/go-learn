package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maximumLength([]int{5, 4, 1, 2, 2}))
}

func maximumLength(nums []int) int {
	cntMap := make(map[int]int)
	for _, num := range nums {
		cntMap[num]++
	}
	values := make([]int, 0, len(cntMap))
	for key := range cntMap {
		values = append(values, key)
	}
	sort.Ints(values)
	result := 1
	if result < cntMap[1] {
		result = cntMap[1]
	}
	if result%2 == 0 {
		result--
	}
	pre := make(map[int]int)
	for _, val := range values {
		sqrt := int(math.Sqrt(float64(val)))
		pre[val] = 1
		if val == 1 || sqrt*sqrt != val {
			continue
		}
		if cntMap[sqrt] >= 2 {
			pre[val] = pre[sqrt] + 1
		}
		tmp := pre[val]*2 - 1
		if tmp > result {
			result = tmp
		}
	}
	return result
}
