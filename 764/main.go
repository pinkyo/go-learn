package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
	fmt.Println(abs(-3))
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

// Binary search function to find the index of a target element in a sorted array
func binarySearch(array []int, target int) int {
	left, right := 0, len(array)-1 // Initialize left and right pointers to the first and last indices of the array
	for left <= right {            // Keep searching until left is greater than right (i.e., not found)
		mid := (left + right) / 2 // Calculate the middle index
		if array[mid] == target { // If the element at the middle index matches the target, return its index
			return mid
		} else if array[mid] < target { // If the element is smaller than the target, move left pointer to the right of the middle index
			left = mid + 1
		} else { // If the element is greater than the target, move right pointer to the left of the middle index
			right = mid - 1
		}
	}
	return left // Return the left index if not found (since it's the smallest possible index that could be the target)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Function to calculate the absolute value of an integer.
func abs(x int) int {
	// Return the negation of x if it is less than 0, otherwise return x itself
	if x < 0 {
		return -1 * x
	}

	return x
}
