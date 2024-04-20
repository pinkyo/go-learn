package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	columnZeroMap := make(map[int][]int)
	rowZeroMap := make(map[int][]int)
	for _, mine := range mines {
		columnZeroMap[mine[1]] = append(columnZeroMap[mine[1]], mine[0])
		rowZeroMap[mine[0]] = append(rowZeroMap[mine[0]], mine[1])
	}
	for column, zeroRows := range columnZeroMap {
		zeroRows = append([]int{-1}, zeroRows...)
		zeroRows = append(zeroRows, n)
		sort.Ints(zeroRows)
		columnZeroMap[column] = zeroRows
	}
	for row, zeroColumns := range rowZeroMap {
		zeroColumns = append([]int{-1}, zeroColumns...)
		zeroColumns = append(zeroColumns, n)
		sort.Ints(zeroColumns)
		rowZeroMap[row] = zeroColumns
	}
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			columnZeros := columnZeroMap[j]
			if columnZeros == nil {
				columnZeros = []int{-1, n}
			}
			rowZeros := rowZeroMap[i]
			if rowZeros == nil {
				rowZeros = []int{-1, n}
			}
			idx1 := binarySearch(columnZeros, i)
			idex2 := binarySearch(rowZeros, j)
			val := min(columnZeros[idx1]-i, i-columnZeros[idx1-1])
			val2 := min(rowZeros[idex2]-j, j-rowZeros[idex2-1])
			tmp := min(val, val2)
			if tmp > result {
				result = tmp
			}
		}
	}

	return result
}

func binarySearch(array []int, target int) int {
	left, right := 0, len(array)-1
	for left <= right {
		mid := (left + right) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
