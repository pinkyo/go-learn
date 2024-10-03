package main

func main() {
	// Call the clearDigits function with the input string
	input := "123abc456def789"
	// Print the result
	println(clearDigits(input))
}

func clearDigits(s string) string {
	// Initialize an empty string to store the result
	result := ""
	// Loop through each character in the string
	for _, c := range s {
		// Use the isDigit function to check if the character is a digit
		if !isDigit(c) {
			// If not a digit, add it to the result string
			result += string(c)
			// If it's a digit, check if the last character in the result string is a non-digit
		} else if len(result) > 0 && !isDigit(rune(result[len(result)-1])) {
			// If the last character is a non-digit, remove it from the result string
			result = result[:len(result)-1]
		}
	}
	// Return the result string
	return result
}

// isDigit checks if a rune is a digit.
func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
