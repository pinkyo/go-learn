package main

func main() {
}

// This function takes an array of tasks and a space as input and returns an int64 as the result
func taskSchedulerII(tasks []int, space int) int64 {
	// Initialize the result variable
	var result int64
	// Create a map to store the previous task's result
	preMap := make(map[int]int64)
	// Iterate through the tasks array
	for _, task := range tasks {
		// If the task is in the map, calculate the next task's result
		if pre, ok := preMap[task]; ok {
			next := pre + int64(space)
			// If the next task's result is greater than the current result, update the result
			if next > result {
				result = next
			}
		}
		// Increment the result by 1
		result++
		// Store the current task's result in the map
		preMap[task] = result
	}
	// Return the result
	return result
}
