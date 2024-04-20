package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("You are not alone.")
    fmt.Println("We are the world.")
    fmt.Println("")
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
