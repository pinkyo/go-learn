package main

import "fmt"

func main() {
	// fmt.Println(countOfSubstrings("aeiou", 0))
	fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1))
}

func countOfSubstrings(word string, k int) int64 {
	n := len(word)
	if n == 0 {
		return 0
	}

	result := int64(0)

	// Sliding window approach
	// left: left boundary of window
	// right: right boundary of window
	left := 0

	// Count of each vowel in current window
	vowelCount := make(map[byte]int)
	// Count of consonants in current window
	consonantCount := 0

	for right := 0; right < n; right++ {
		// Add current character to window
		if isVowel(word[right]) {
			vowelCount[word[right]]++
		} else {
			consonantCount++
		}

		// Shrink window from left while conditions are violated
		// We need exactly 5 distinct vowels and at most k consonants
		for consonantCount > k || len(vowelCount) > 5 {
			if isVowel(word[left]) {
				vowelCount[word[left]]--
				if vowelCount[word[left]] == 0 {
					delete(vowelCount, word[left])
				}
			} else {
				consonantCount--
			}
			left++
		}

		// If we have exactly 5 distinct vowels and at most k consonants
		if len(vowelCount) == 5 && consonantCount <= k {
			// Count valid substrings ending at position 'right'
			// We need to find how many valid starting positions exist
			tempLeft := left
			tempVowelCount := make(map[byte]int)
			for k, v := range vowelCount {
				tempVowelCount[k] = v
			}
			tempConsonantCount := consonantCount

			// Move tempLeft to find the rightmost position where we still have all 5 vowels
			for tempLeft <= right {
				if len(tempVowelCount) == 5 && tempConsonantCount <= k {
					result += int64(tempLeft - left + 1)
					break
				}
				if isVowel(word[tempLeft]) {
					tempVowelCount[word[tempLeft]]--
					if tempVowelCount[word[tempLeft]] == 0 {
						delete(tempVowelCount, word[tempLeft])
					}
				} else {
					tempConsonantCount--
				}
				tempLeft++
			}
		}
	}

	return result
}

// Function to check if a given character is a vowel
func isVowel(c byte) bool {
	// Check if the character 'c' is a vowel (a, e, i, o, or u)
	return c == 'a' || // Check if c is 'a'
		c == 'e' || // Check if c is 'e'
		c == 'i' || // Check if c is 'i'
		c == 'o' || // Check if c is 'o'
		c == 'u' // Check if c is 'u'
}
