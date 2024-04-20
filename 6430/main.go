package main

import "fmt"

func main() {
	fmt.Println(circularGameLosers(5, 2))
}

func circularGameLosers(n int, k int) []int {
	countMap := make(map[int]int)
	cur := 0
	var cnt = 1
	for countMap[cur] < 1 {
		countMap[cur]++
		cur = (cur + k*cnt) % n
		cnt++
	}
	var result []int
	for i := 0; i < n; i++ {
		if countMap[i] < 1 {
			result = append(result, i+1)
		}
	}
	return result
}
