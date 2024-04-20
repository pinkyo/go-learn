package main

import "fmt"

func main() {
	// fmt.Println(punishmentNumber(10))
	// fmt.Println(isPunishmentNumber(9))
	fmt.Println(isPunishmentNumber(36))
}

func punishmentNumber(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		if isPunishmentNumber(i) {
			result += i * i
		}
	}
	return result
}

func isPunishmentNumber(n int) bool {
	square := n * n
	var values []int
	for square > 0 {
		values = append(values, square%10)
		square /= 10
	}
	// reverse values
	for i := 0; i < len(values)/2; i++ {
		values[i], values[len(values)-1-i] = values[len(values)-1-i], values[i]
	}
	var sumValues [][]int
	var value int
	for i := 0; i < len(values); i++ {
		value *= 10
		value += values[i]
		sumValues = append(sumValues, []int{value})
	}
	for i := 1; i < len(values); i++ {
		var cur int
		for j := i; j < len(values); j++ {
			cur *= 10
			cur += values[j]
			for _, sum := range sumValues[i-1] {
				temp := sum + cur
				if temp <= n {
					sumValues[j] = append(sumValues[j], temp)
				}
			}
		}
	}
	for _, sum := range sumValues[len(values)-1] {
		if sum == n {
			return true
		}
	}
	return false
}
