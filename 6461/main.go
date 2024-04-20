package main

func main() {

}

func isFascinating(n int) bool {
	digitMap := make(map[int]bool)
	values := []int{n, n * 2, n * 3}
	for _, value := range values {
		for value > 0 {
			digit := value % 10
			if digitMap[digit] {
				return false
			}
			if digit == 0 {
				return false
			}
			digitMap[digit] = true
			value /= 10
		}
	}
	return len(digitMap) == 9
}
