package main

import "fmt"

func main() {
	fmt.Println("Percentage of 'a' in 'banana':", percentageLetter("banana", 'a'))
}

func percentageLetter(s string, letter byte) int {
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			count++
		}
	}
	return count * 100 / len(s)
}
