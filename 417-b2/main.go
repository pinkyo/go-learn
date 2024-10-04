package main

import (
	"fmt"
)

func main() {
	// fmt.Println(countOfSubstrings("aeiou", 0))
	fmt.Println(countOfSubstrings("ieaouqqieaouqq", 1))
}

func countOfSubstrings(word string, k int) int {
	result := 0
	for i := 0; i < len(word); i++ {
		vowel := make(map[byte]int)
		consonants := 0
		for j := i; j < len(word); j++ {
			if isVowel(word[j]) {
				vowel[word[j]]++
			} else {
				consonants++
			}
			if consonants > k {
				break
			}
			if consonants == k && len(vowel) == 5 {
				result++
			}
		}
	}
	return result
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
