package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumBeauty([][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}))
}

func maximumBeauty(items [][]int, queries []int) []int {
	arr := make([]int, len(queries))
	copy(arr, queries)
	sort.Ints(arr)
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] > items[j][1]
		}
		return items[i][0] < items[j][0]
	})
	pos := 0
	resultMap := make(map[int]int)
	curMax := 0
	for i := 0; i < len(arr); i++ {
		for pos < len(items) && items[pos][0] <= arr[i] {
			if items[pos][1] > curMax {
				curMax = items[pos][1]
			}
			pos++
		}
		resultMap[arr[i]] = curMax
	}

	result := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		result[i] = resultMap[queries[i]]
	}

	return result
}
