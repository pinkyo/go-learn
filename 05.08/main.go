package main

import (
	"fmt"
)

func main() {
	// fmt.Println(drawLine(1, 32, 30, 31, 0))
	// fmt.Println(drawLine(3, 96, 0, 95, 0))
	fmt.Println(quickSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 6}))
}

func drawLine(length int, w int, x1 int, x2 int, y int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		var cur int32
		from := 32 * i
		if from < x1 {
			from = x1
		}
		to := 32 * (i + 1)
		if to > x2+1 {
			to = x2 + 1
		}
		for j := from; j < to; j++ {
			cur |= 1 << uint(31-j%32)
		}
		// fmt.Printf("%b", cur)
		result[i] = int(cur)
	}

	return result
}

// quick sort algorithm
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	less := make([]int, 0)
	greater := make([]int, 0)
	for _, v := range arr[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}
	leftResult := quickSort(less)
	rightResult := quickSort(greater)

	result := append(leftResult, pivot)
	result = append(result, rightResult...)
	return result
}
