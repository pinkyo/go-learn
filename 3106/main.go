package main

func main() {
}

// Function to get the smallest string from a given input string and k
func getSmallestString(s string, k int) string {
	// Initialize result string with empty value
	result := ""

	// Iterate through each character in the input string
	for i := range s {
		// Calculate the distance of current character from 'a' and 26 minus that distance
		dis := min(int(s[i]-'a'), int(26-s[i]+'a'))

		// Check if the calculated distance is less than or equal to k
		if dis <= k {
			// Add 'a' to result string and decrease k by the distance value
			result += "a"
			k -= dis
		} else {
			// Add current character to result string with k as the offset, then append the rest of the input string
			result += string(s[i] - byte(k))
			result += s[i+1:]
			break // Break loop if the condition is met, since we've already added the required characters
		}
	}

	// Return the final result string
	return result
}

// Helper function to find the minimum of two integers
func min(any ...int) int {
	result := any[0]
	for i := 1; i < len(any); i++ {
		if result > any[i] {
			result = any[i]
		}
	}

	return result
}
