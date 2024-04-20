package main

func main() {
}

func countWays(nums []int) int {
	countArr := make([]int, len(nums)+1)
	for _, num := range nums {
		countArr[num]++
	}

	var result int
	if countArr[0] == 0 {
		result++
	}
	for i := 1; i <= len(countArr); i++ {
		cur := countArr[i]
		countArr[i] += countArr[i-1]
		if countArr[i] == i && cur == 0 {
			result++
		}
	}
	return result
}
