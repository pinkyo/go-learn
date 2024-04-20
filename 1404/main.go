package main

func main() {
}

func numSteps(s string) int {
	var result int
	var increment bool
	for i := len(s) - 1; i >= 0; i-- {
		if i == 0 && !increment {
			break
		}
		if s[i] == '0' && !increment || s[i] == '1' && increment {
			result++
			continue
		}
		result += 2
		increment = true
	}
	return result
}
