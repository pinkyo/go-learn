package main

func main() {
}

func commonFactors(a int, b int) int {
	limit := a
	if b < a {
		limit = b
	}
	var result int
	for i := 1; i <= limit; i++ {
		if a%i == 0 && b%i == 0 {
			result++
		}
	}
	return result
}
