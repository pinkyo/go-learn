package main

func main() {
}

// This function returns the kth character in the infinite string "abcdefghijklmnopqrstuvwxyz"
func kthCharacter(k int) byte {
	// Initialize the result as "a"
	result := "a"
	// Loop until the length of the result is greater than or equal to k
	for len(result) < k {
		// Initialize the next string as an empty string
		next := ""
		// Loop through each character in the result
		for _, c := range result {
			// Append the next character to the next string
			next += string(c + 1)
		}
		// Append the next string to the result
		result += next
	}

	// Return the kth character in the result
	return result[k-1]
}
