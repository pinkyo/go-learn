package main

func main() {
}

func sumOfSquares(nums []int) int {
	result := 0
	for i, num := range nums {
		if len(nums)%(i+1) == 0 {
			result += num * num
		}
	}
	return result
}
