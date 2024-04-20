package main

import "fmt"

func main() {
	// fmt.Println(int('a' - 'A'))
	fmt.Println(countKeyChanges("aAbBcC"))
}

func countKeyChanges(s string) int {
	sameCase := []int{0, 'a' - 'A'}

	result := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] || abs(int(s[i])-int(s[i-1])) == sameCase[1] {
			continue
		}
		result++
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
