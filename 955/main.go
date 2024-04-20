package main

import "fmt"

func main() {
	fmt.Println(minDeletionSize([]string{"xga", "xfb", "yfa"}))
}

func minDeletionSize(strs []string) int {
	var result int
	prefix := make([]string, len(strs))
OUTER:
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if prefix[j]+string(strs[j][i]) < prefix[j-1]+string(strs[j-1][i]) {
				result++
				continue OUTER
			}
		}
		for j := 0; j < len(strs); j++ {
			prefix[j] += string(strs[j][i])
		}
	}
	return result
}
