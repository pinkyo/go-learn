package main

import "fmt"

func main() {
	fmt.Println(maximumPrimeDifference([]int{4, 2, 9, 5, 3}))
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func maximumPrimeDifference(nums []int) int {
	left, right := 0, len(nums)-1
	for left < len(nums) && !isPrime(nums[left]) {
		left++
	}
	for right >= 0 && !isPrime(nums[right]) {
		right--
	}
	return right - left
}
