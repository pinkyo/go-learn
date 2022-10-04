package main

import "fmt"

func main() {
	// for i := 4; i < 1000000; i++ {
	// 	if isStrictlyPalindromic(i) {
	// 		println(i)
	// 	}
	// }
	// println(isPalindrome(9, 2))
	// fmt.Println(maximumRows([][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}, 2))
	fmt.Println(maximumRows([][]int{
		{1, 0, 1, 1, 0, 0},
		{1, 1, 1, 0, 1, 0},
		{0, 0, 1, 1, 1, 1},
		{1, 0, 1, 1, 1, 0},
		{1, 0, 0, 1, 0, 1},
		{0, 0, 1, 1, 0, 0},
		{0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 0, 1},
	}, 5))
}

func maximumRows(mat [][]int, cols int) int {
	// rowCnt := make([]int, len(mat[0]))
	// for i := 0; i < len(mat); i++ {
	// 	for j := 0; j < len(mat[i]); j++ {
	// 		if mat[i][j] == 1 {
	// 			rowCnt[j]++
	// 		}
	// 	}
	// }
	selected := make([]bool, len(mat[0]))
	count := func() int {
		var cnt int
	OUTER:
		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[i]); j++ {
				if !selected[j] && mat[i][j] == 1 {
					continue OUTER
				}
			}
			cnt++
		}
		return cnt
	}
	var max int
	var backtrace func(int, int)
	backtrace = func(idx, cnt int) {
		if cnt == 0 || len(mat[0]) == idx+cnt {
			for i := 0; i < cnt; i++ {
				selected[idx+i] = true
			}
			val := count()
			if val > max {
				max = val
			}
			for i := 0; i < cnt; i++ {
				selected[idx+i] = false
			}
			return
		}
		// if rowCnt[idx] < len(mat[0])-max {
		backtrace(idx+1, cnt)
		// }
		selected[idx] = true
		backtrace(idx+1, cnt-1)
		selected[idx] = false
	}
	backtrace(0, cols)
	return max
}
