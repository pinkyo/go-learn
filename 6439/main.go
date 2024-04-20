package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minLength("ABFCACDB"))
}

func minLength(s string) int {
	result := s
	for {
		idx1 := strings.Index(result, "AB")
		if idx1 >= 0 {
			result = result[:idx1] + result[idx1+2:]
			continue
		}
		idx2 := strings.Index(result, "CD")
		if idx2 >= 0 {
			result = result[:idx2] + result[idx2+2:]
			continue
		}
		break
	}
	return len(result)
}
