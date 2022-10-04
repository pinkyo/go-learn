package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumMoney([][]int{
		{2, 1},
		{5, 0},
		{4, 2},
	}))
}

func minimumMoney(transactions [][]int) int64 {
	var subArr [][]int
	var max int
	for i := 0; i < len(transactions); i++ {
		if transactions[i][1]-transactions[i][0] < 0 {
			subArr = append(subArr, transactions[i])
		} else if transactions[i][0] > max {
			max = transactions[i][0]
		}
	}
	sort.Sort(IntPairSlice(subArr))
	var result int64
	var cur int
	for _, v := range subArr {
		if cur < v[0] {
			result += int64(v[0] - cur)
			cur = v[0]
		}
		cur += v[1] - v[0]
	}
	if cur < max {
		result += int64(max - cur)
	}
	return result
}

type IntPairSlice [][]int

func (a IntPairSlice) Len() int      { return len(a) }
func (a IntPairSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a IntPairSlice) Less(i, j int) bool {
	val1 := a[i][1] - a[i][0]
	val2 := a[j][1] - a[j][0]
	if val1 == val2 {
		return a[i][0] < a[j][0]
	}
	return val1 < val2
}
