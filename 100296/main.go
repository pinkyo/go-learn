package main

func main() {
}

func findPermutationDifference(s string, t string) int {
	rune2index := make(map[rune]int)
	for i, c := range s {
		rune2index[c] = i
	}
	diff := 0
	for i, c := range t {
		diff += abs(i - rune2index[c])
	}
	return diff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
