package main

import "fmt"

func main() {
	fmt.Println(makeSmallestPalindrome("egcfe"))
	fmt.Println(makeSmallestPalindrome("abcd"))
	fmt.Println(makeSmallestPalindrome("seven"))
}

func makeSmallestPalindrome(s string) string {
	var result string
	for i := 0; i < len(s)/2; i++ {
		if s[i] < s[len(s)-1-i] {
			result += string(s[i])
		} else {
			result += string(s[len(s)-1-i])
		}
	}
	var reverse string
	for i := len(result) - 1; i >= 0; i-- {
		reverse += string(result[i])
	}
	if len(s)%2 == 1 {
		result += string(s[len(s)/2])
	}
	return result + reverse
}
