package main

func main() {
}

func minChanges(s string) int {
	result := 0
	for i := 1; i < len(s); i += 2 {
		if s[i] != s[i-1] {
			result++
		}
	}
	return result
}
