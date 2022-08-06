package main

import "fmt"

func main() {
	fmt.Println(numRabbits([]int{1, 0, 1, 0, 0}))
	fmt.Println(numRabbits([]int{10, 10, 10}))
}

func numRabbits(answers []int) int {
	cntMap := make(map[int]int)
	for _, v := range answers {
		cntMap[v]++
	}
	var result int
	for k, v := range cntMap {
		if v <= k+1 {
			result += k + 1
			continue
		}
		reminder := v % (k + 1)
		result += v - reminder
		if reminder > 0 {
			result += (k + 1)
		}
	}
	return result
}
