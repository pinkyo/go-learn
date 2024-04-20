package main

import "fmt"

func main() {
	fmt.Println(countBlackBlocks(3, 3, [][]int{
		{0, 0}, {1, 1}, {0, 2},
	}))
}

func countBlackBlocks(m int, n int, coordinates [][]int) []int64 {
	row2columns := make(map[int][]int)
	for _, coordinate := range coordinates {
		row2columns[coordinate[0]] = append(row2columns[coordinate[0]], coordinate[1])
	}
	result := make([]int64, 5)
	for i := 0; i < m-1; i++ {
		values := row2columns[i]
		values = append(values, row2columns[i+1]...)
		cntMap := make(map[int]int)
		for _, value := range values {
			cntMap[value]++
		}
		black := 0
		for col, cnt := range cntMap {
			if col > 0 {
				result[cnt+cntMap[col-1]]++
				black++
			}
			if col+1 < n {
				if cnt2, ok := cntMap[col+1]; !ok {
					result[cnt+cnt2]++
					black++
				}
			}
		}
		result[0] += int64(n - 1 - black)
	}
	return result
}
