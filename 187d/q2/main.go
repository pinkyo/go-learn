package main

import "fmt"

func main() {
	fmt.Println(maximumValue(3, 7, 7))
}

func maximumValue(n int, s int, m int) int64 {
	if n == 1 {
		return int64(s)
	}
	result := s + n/2*m - (n-1)/2
	if n%2 != 0 {
		result += 1
	}
	return int64(result)
}
