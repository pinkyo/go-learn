package main

import "fmt"

func main() {
	fmt.Println(matrixSumQueries(8,
		[][]int{{0, 6, 30094}, {0, 7, 99382}, {1, 2, 18599}, {1, 3, 49292}, {1, 0, 81549}, {1, 1, 38280}, {0, 0, 19405}, {0, 4, 30065}, {1, 4, 60826}, {1, 5, 9241}, {0, 5, 33729}, {0, 1, 41456}, {0, 2, 62692}, {0, 3, 30807}, {1, 7, 70613}, {1, 6, 9506}, {0, 5, 39344}, {1, 0, 44658}, {1, 1, 56485}, {1, 2, 48112}, {0, 6, 43384}}))
}

func matrixSumQueries(n int, queries [][]int) int64 {
	rowSet := make(map[int]struct{})
	colSet := make(map[int]struct{})
	var result int64
	for i := len(queries) - 1; i >= 0; i-- {
		query := queries[i]
		if _, ok := rowSet[query[1]]; query[0] == 0 && !ok {
			result += int64(query[2] * (n - len(colSet)))
			rowSet[query[1]] = struct{}{}
		}
		if _, ok := colSet[query[1]]; query[0] == 1 && !ok {
			result += int64(query[2] * (n - len(rowSet)))
			colSet[query[1]] = struct{}{}
		}
	}
	return result
}
