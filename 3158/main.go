package main

func main() {
}

// Function to find duplicate numbers using XOR operation.
func duplicateNumbersXOR(nums []int) int {
	// Initialize a map to count the occurrences of each number in the array.
	cntMap := make(map[int]int)

	// Iterate through each number in the array.
	for _, num := range nums {
		// Increase the count for the current number if it exists, otherwise add a new entry with a count of 1.
		cntMap[num]++
	}

	// Initialize a variable to store the final result (the XOR of the duplicated numbers).
	result := 0

	// Iterate through each key-value pair in the map, where the key is a number and the value is its count.
	for val, cnt := range cntMap {
		// If the current number occurs only once, skip it.
		if cnt == 1 {
			continue
		}
		// Calculate the XOR of the result with the current number.
		result ^= val
	}

	// Return the calculated result, which should be the duplicated numbers XOR'd together.
	return result
}
