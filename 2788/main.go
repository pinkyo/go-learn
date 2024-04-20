package main

import "strings"

func main() {
}

func splitWordsBySeparator(words []string, separator byte) []string {
	var result []string
	for _, word := range words {
		ws := strings.Split(word, string(separator))
		for _, w := range ws {
			if w != "" {
				result = append(result, w)
			}
		}
	}
	return result
}
