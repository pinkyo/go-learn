package main

import "fmt"

func main() {
	fmt.Println(isValid("aabcbabcc"))
}

func isValid(s string) bool {
	bs := make([]byte, len(s))
	var top int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			bs[top] = 'a'
			top++
		case 'b':
			if top == 0 || bs[top-1] != 'a' {
				return false
			}
			bs[top] = 'b'
			top++
		case 'c':
			if top < 2 || bs[top-1] != 'b' || bs[top-2] != 'a' {
				return false
			}
			top -= 2
		default:
			return false
		}
	}
	return top == 0
}
