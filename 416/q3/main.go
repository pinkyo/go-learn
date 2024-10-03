package main

import "fmt"

func main() {
	fmt.Println(validSubstringCount("abcabc", "abc"))
}

func validSubstringCount(word1 string, word2 string) int64 {
	// Initialize an array to count the occurrences of each character in word2
	cntArr := [26]int64{0}
	for _, cn := range word2 {
		cntArr[cn-'a']++
	}

	// Initialize an array to keep track of the current occurrences of each character in word1
	curArr := [26]int64{0}

	// Define a helper function to check if the current substring is valid
	isValid := func() bool {
		for i := 0; i < 26; i++ {
			if cntArr[i] > curArr[i] {
				// If any character in word2 has more occurrences than in the current substring, return false
				return false
			}
		}
		// If all characters in word2 have equal or fewer occurrences in the current substring, return true
		return true
	}

	result := 0
	pre := 0

	// Iterate over each character in word1
	for idx, ch := range word1 {
		// Increment the count of the current character in curArr
		curArr[ch-'a']++

		// Check if the current substring is valid
		if isValid() {
			// If valid, add the number of substrings ending at the current index to the result
			result += len(word1) - idx
		} else {
			// If not valid, continue to the next iteration
			continue
		}

		// Start from the previous index and shrink the substring by decrementing the count of each character
		for pre <= idx {
			curArr[word1[pre]-'a']--
			pre++

			// Check if the shrunk substring is valid
			if isValid() {
				// If valid, add the number of substrings ending at the current index to the result
				result += len(word1) - idx
			} else {
				// If not valid, break out of the loop
				break
			}
		}
	}

	// Return the final result as an int64
	return int64(result)
}
