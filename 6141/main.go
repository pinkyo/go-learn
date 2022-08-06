package main

import "sort"

func main() {
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	valueWeights := make(map[int]int)
	for _, item := range items1 {
		valueWeights[item[0]] += item[1]
	}
	for _, item := range items2 {
		valueWeights[item[0]] += item[1]
	}
	var result [][]int
	for k, v := range valueWeights {
		result = append(result, []int{k, v})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}
