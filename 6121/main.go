package main

import "sort"

func main() {

}

type Item struct {
	val string
	idx int
}

type ItemSlice []Item

func (this ItemSlice) Less(i, j int) bool {
	if this[i].val == this[j].val {
		return this[i].idx < this[j].idx
	}

	return this[i].val < this[j].val
}

func (this ItemSlice) Len() int {
	return len(this)
}

func (this ItemSlice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	strMatrix := make([][]Item, len(nums[0]))
	for j := 0; j < len(nums[0]); j++ {
		strMatrix[j] = make([]Item, len(nums))
		for i, v := range nums {
			strMatrix[j][i] = Item{v[len(v)-j-1:], i}
		}
	}
	for i := 0; i < len(strMatrix); i++ {
		sort.Stable(ItemSlice(strMatrix[i]))
	}
	res := make([]int, len(queries))
	for i, query := range queries {
		res[i] = strMatrix[query[1]-1][query[0]-1].idx
	}
	return res
}
