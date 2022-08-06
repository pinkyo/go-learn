package main

import "fmt"

func main() {
	fmt.Println(removeAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"}))
}

func removeAnagrams(words []string) []string {
	var result []string
	result = append(result, words[0])
	for i := 1; i < len(words); i++ {
		if isAnagram(words[i], words[i-1]) {
			continue
		}
		result = append(result, words[i])
	}
	return result
}

func isAnagram(word string, word2 string) bool {
	if len(word) != len(word2) {
		return false
	}
	cntArr := make([]int, 26)
	for i := 0; i < len(word); i++ {
		cntArr[word[i]-'a']++
	}
	for i := 0; i < len(word2); i++ {
		cntArr[word2[i]-'a']--
	}
	for i := 0; i < len(cntArr); i++ {
		if cntArr[i] != 0 {
			return false
		}
	}
	return true
}
