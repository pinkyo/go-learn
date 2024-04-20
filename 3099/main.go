package main

func main() {
}

func sumOfTheDigitsOfHarshadNumber(x int) int {
	sum := 0
	t := x
	for t > 0 {
		sum += t % 10
		t /= 10
	}
	if x%sum == 0 {
		return sum
	} else {
		return -1
	}
}
