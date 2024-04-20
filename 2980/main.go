package main

func main() {
}

func hasTrailingZeros(nums []int) bool {
	count := 0
	for _, num := range nums {
		if num&1 == 0 { // 检查num的最后一位是否为0
			count++
		}
	}
	return count >= 2
}
