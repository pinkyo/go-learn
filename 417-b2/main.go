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
	consonants := 0
	cntMap := make(map[rune]int)
	pre := 0
	pre, i := 0, 0
	for i < len(word) {
		c := rune(word[i])
		if isVowel(c) {
			cntMap[c]++
		} else {
			consonants++
		}
		if consonants == k && len(cntMap) == 5 {
			for pre <= i {
				c := rune(word[pre])
				if isVowel(c) {
					cntMap[c]--
					if cntMap[c] == 0 {
						delete(cntMap, c)
					}
				} else {
					break
				}
				pre++
			}
		}
		for pre <= i && consonants > k {
			c := rune(word[pre])
			if isVowel(c) {
				cntMap[c]--
				if cntMap[c] == 0 {
					delete(cntMap, c)
				}
			} else {
				consonants--
			}
			pre++
		}
	}

	return result
}

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
