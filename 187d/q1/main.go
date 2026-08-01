package main

import "strings"

func main() {
}

func rearrangeString(s string, x byte, y byte) string {
	yc := 0
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == y {
			yc++
		} else {
			result += string(s[i])
		}
	}
	if y > 0 {
		result = strings.Repeat(string(y), yc) + result
	}
	return result
}
