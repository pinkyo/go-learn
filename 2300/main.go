package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	var result []int
	for _, v := range spells {
		target := (int(success)-1)/v + 1
		l, h := 0, len(potions)-1
		for l <= h {
			mid := (l + h) / 2
			if potions[mid] >= target {
				h = mid - 1
			} else {
				l = mid + 1
			}
		}
		result = append(result, len(potions)-l)
	}
	return result
}
