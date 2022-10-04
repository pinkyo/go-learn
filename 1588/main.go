package main

import "fmt"

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
}

func sumOddLengthSubarrays(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		var temp int
		for j := i; j < len(arr); j++ {
			temp += arr[j]
			if (j-i+1)%2 == 1 {
				sum += temp
			}
		}
	}
	return sum
}
