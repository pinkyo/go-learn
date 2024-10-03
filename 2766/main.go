package main

import "sort"

func main() {
}

// relocateMarbles Function to relocate marbles
func relocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
	// Create a map to store the number and its corresponding count
	numMap := make(map[int]int)
	// Traverse the nums slice and record the number of occurrences of each number in the map
	for _, num := range nums {
		numMap[num]++
	}
	// Traverse the moveFrom slice
	for i := 0; i < len(moveFrom); i++ {
		// If the source and destination positions are the same, skip this operation
		if moveFrom[i] == moveTo[i] {
			continue
		}
		// Update the count of the number at the destination position in the map
		numMap[moveTo[i]] += numMap[moveFrom[i]]
		// Set the count of the number at the source position to 0
		numMap[moveFrom[i]] = 0
	}
	// Create a slice to store the final result
	result := make([]int, 0, len(nums))
	// Traverse the numMap map
	for k, v := range numMap {
		// If the number occurs more than once, add it to the result slice
		if v > 0 {
			result = append(result, k)
		}
	}
	// The result slice is sorted
	sort.Ints(result)
	// Returns the sorted result slice
	return result
}
