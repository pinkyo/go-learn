package main

import "fmt"

func main() {
	// fmt.Println(minInitialStrength([]int{5, 10, 15}, [][]int{{1, 2, 10}, {1, 2, 5}}))
	fmt.Println(minInitialStrength([]int{387}, [][]int{{0, 0, 361}}))
}

func minInitialStrength(monsters []int, boosts [][]int) int64 {
	n := len(monsters)

	diff := make([]int, n+1)
	for _, boost := range boosts {
		l, r, v := boost[0], boost[1], boost[2]
		diff[l] += v
		diff[r+1] -= v
	}

	mm := make([]int, n)
	bonus := 0
	for i := 0; i < n; i++ {
		bonus += diff[i]
		mm[i] = monsters[i] - bonus
	}

	norvelithx := monsters

	ri := n - 1
	for ri >= 0 && mm[ri] <= 0 {
		ri--
	}
	var result int64 = 0
	for i := 0; i < ri; i++ {
		result += int64(norvelithx[i])
	}
	if ri >= 0 {
		result += int64(mm[ri])
	}
	return result
}