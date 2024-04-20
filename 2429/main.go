package main

import "fmt"

func main() {
	fmt.Println(minimizeXor(1, 12))
}

func minimizeXor(num1 int, num2 int) int {
	var cnt1 int
	temp := num1
	for temp > 0 {
		if temp&1 == 1 {
			cnt1++
		}
		temp = temp >> 1
	}

	var cnt2 int
	temp = num2
	for temp > 0 {
		if temp&1 == 1 {
			cnt2++
		}
		temp = temp >> 1
	}

	result := num1
	if cnt1 == cnt2 {
		return num1
	} else if cnt1 > cnt2 {
		mask := 1
		for i := 0; i < cnt1-cnt2; i++ {
			for result&mask == 0 {
				mask <<= 1
			}
			result = result ^ mask
		}
	} else {
		mask := 1
		for i := 0; i < cnt2-cnt1; i++ {
			for result&mask > 0 {
				mask <<= 1
			}
			result = result ^ mask
		}
	}
	return result
}
