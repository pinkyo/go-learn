package main

func main() {
}

func sumIndicesWithKSetBits(nums []int, k int) int {
	var result int
	for i, num := range nums {
		if countBinaryDigit1(i) == k {
			result += num
		}
	}
	return result
}

func countBinaryDigit1(value int) int {
	count := 0
	for value > 0 {
		value &= value - 1 // reset the right-most 1-bit of value
		count++
	}
	return count
}
