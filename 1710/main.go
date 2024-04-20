package main

import "sort"

func main() {
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Sort(boxSlice(boxTypes))
	var result int
	for i := 0; i < len(boxTypes); i++ {
		if boxTypes[i][0] <= truckSize {
			result += boxTypes[i][0] * boxTypes[i][1]
			truckSize -= boxTypes[i][0]
		} else {
			result += truckSize * boxTypes[i][1]
			break
		}
	}
	return result
}

type boxSlice [][]int

func (b boxSlice) Len() int {
	return len(b)
}

func (b boxSlice) Less(i, j int) bool {
	return b[i][1] > b[j][1]
}

func (b boxSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
