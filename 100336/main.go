package main

func main() {
}

// numberOfAlternatingGroups function counts the number of groups in the given slice of colors where
// each color is different from its neighbors.
func numberOfAlternatingGroups(colors []int) int {
	// Initialize the result variable to store the count of alternating groups
	result := 0

	// Iterate through the colors slice
	for i := 0; i < len(colors); i++ {
		// Get the previous color by taking the modulo of the index minus one with the length of the slice
		pre := colors[(i-1+len(colors))%len(colors)]

		// Get the next color by taking the modulo of the index plus one with the length of the slice
		next := colors[(i+1)%len(colors)]

		// Check if the current color is different from both the previous and next colors
		if pre != colors[i] && next != colors[i] {
			// If it is, increment the result as we found an alternating group
			result++
		}
	}

	// Return the final count of alternating groups
	return result
}
