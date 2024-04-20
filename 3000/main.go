package main

func main() {
}

func areaOfMaxDiagonal(dimensions [][]int) int {
	max := 0
	result := 0

	for _, dimension := range dimensions {
		val := dimension[0]*dimension[0] +
			dimension[1]*dimension[1]
		if val < max {
			continue
		}
		area := dimension[0] * dimension[1]
		if val == max && area > result || val > max {
			max = val
			result = area
		}
	}

	return result
}
