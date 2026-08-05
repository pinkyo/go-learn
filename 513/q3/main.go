package main

import "sort"

func main() {
	// countTasks([]int{1, 4, 4}, []int{9, 1, 4})
	countTasks([]int{2, 3, 4}, []int{20, 4, 5})
}

func countTasks(tasks []int, shifts []int) []int {
	prefixSum := make([]int, len(tasks)+1)
	for i := 0; i < len(tasks); i++ {
		prefixSum[i+1] = prefixSum[i] + tasks[i]
	}
	result := make([]int, len(shifts))
	sum := 0
	for i := 0; i < len(shifts); i++ {
		next := shifts[i] + sum
		if next >= prefixSum[len(prefixSum)-1] {
			result[i] = 0
			sum = 0
			continue
		}
		j := sort.Search(len(prefixSum), func(j int) bool {
			return prefixSum[j] > next
		})
		if j == -1 {
			result[i] = len(tasks)
			sum = 0
			continue
		}
		result[i] = len(prefixSum) - j
		sum = next
	}
	return result
}
