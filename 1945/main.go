package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getLucky("leetcode", 2))
}

func getLucky(s string, k int) int {
	var str string
	for i := 0; i < len(s); i++ {
		str += strconv.Itoa(int(s[i] - 'a' + 1))
	}
	var result int
	for i := 0; i < k; i++ {
		result = 0
		for j := 0; j < len(str); j++ {
			result += int(str[j] - '0')
		}
		str = strconv.Itoa(result)
	}
	return result
}
