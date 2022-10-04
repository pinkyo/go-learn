package main

func main() {
}

func countBalls(lowLimit int, highLimit int) int {
	box := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		box[sum(i)]++
	}
	var max int
	for _, v := range box {
		if v > max {
			max = v
		}
	}
	return max
}

func sum(n int) int {
	var sum int
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
