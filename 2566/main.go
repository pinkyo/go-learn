package main

import "fmt"

func main() {
	fmt.Println(minMaxDifference(1234))
}

func minMaxDifference(num int) int {
	return getMaxVal(num) - getMinVal(num)
}

func getMaxVal(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	digit := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digit = digits[i]
			break
		}
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == digit {
			digits[i] = 9
		}
	}
	result := 0
	for i := len(digits) - 1; i >= 0; i-- {
		result = result*10 + digits[i]
	}
	return result
}

func getMinVal(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num = num / 10
	}
	digit := digits[len(digits)-1]
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == digit {
			digits[i] = 0
		}
	}
	result := 0
	for i := len(digits) - 1; i >= 0; i-- {
		result = result*10 + digits[i]
	}
	return result
}
