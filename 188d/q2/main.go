package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumWidth([]int{1, 3, 2, 5, 7, 5, 4, 2, 1}))
	// fmt.Println(maximumWidth([]int{59, 37, 55, 90, 94}))
}

func maximumWidth(planks []int) int {
	plankCntMap := make(map[int]int, len(planks))
	result := 0
	for i := 0; i < len(planks); i++ {
		plankCntMap[planks[i]]++
	}
	uniqPlanks := make([]int, 0, len(planks))
	for height := range plankCntMap {
		uniqPlanks = append(uniqPlanks, height)
	}
	sort.Ints(uniqPlanks)

	plankSumMap := make(map[int]int)
	for _, height := range uniqPlanks {
		plankSumMap[height] += plankCntMap[height]
		plankSumMap[height*2] += plankCntMap[height] / 2
		for _, height2 := range uniqPlanks {
			if height2 > height {
				break
			}
			plankSumMap[height+height2] += min(plankCntMap[height], plankCntMap[height2])
		}
	}

	for _, cnt := range plankSumMap {
		result = max(result, cnt)
	}
	return result
}
