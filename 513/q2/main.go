package main

func main() {
}

func countRatioSubarrays(nums []int, a int, b int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		oddCnt, evenCnt := 0, 0
		for j := i; j < len(nums); j++ {
			if nums[j]%2 == 0 {
				evenCnt++
			} else {
				oddCnt++
			}
			if oddCnt > 0 && evenCnt*b <= oddCnt*a {
				result++
			}
		}
	}
	return result
}
