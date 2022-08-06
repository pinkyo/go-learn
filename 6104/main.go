package main

import "strings"

func main() {

}

func countAsterisks(s string) int {
	strs := strings.Split(s, "|")
	var result int
	for i := 0; i < len(strs); i++ {
		if i%2 == 1 {
			continue
		}
		for j := 0; j < len(strs[i]); j++ {
			if strs[i][j] == '*' {
				result++
			}
		}
	}
	return result
}
