package main

import "fmt"

func main() {
	fmt.Println(checkDistances(
		"abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	))
	fmt.Println(checkDistances(
		"aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	))
}

func checkDistances(s string, distance []int) bool {
	preMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if pre, ok := preMap[s[i]]; ok {
			if i-pre-1 != distance[s[i]-'a'] {
				return false
			}
		}
		preMap[s[i]] = i
	}
	return true
}
