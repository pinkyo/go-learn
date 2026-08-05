package main

func main() {
}

func maxDigitRange(nums []int) int {
	result := 0
	max := 0

	for _, num := range nums {
		dr := digitRange(num)
		if dr > max {
			result = num
			max = dr
		} else if dr == max {
			result += num
		}
	}
	return result
}

func digitRange(num int) int {
	max, min := 0, 9
	for num > 0 {
		digit := num % 10
		if digit > max {
			max = digit
		}
		if digit < min {
			min = digit
		}
		num /= 10
	}
	return max - min
}
