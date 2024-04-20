package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}

func singleNumber(nums []int) int {
	val1 := 0
	val2 := 0
	for _, num := range nums {
		tmp := (val1 & num)
		val1 ^= num
		val1 ^= (tmp & val2)
		val2 ^= tmp
	}
	return val2 ^ val1
}
