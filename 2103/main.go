package main

func main() {
}

func countPoints(rings string) int {
	state := make([]int, 10)
	for i := 0; i < len(rings)-1; i += 2 {
		val := 0
		switch rings[i] {
		case 'R':
			val = 0
		case 'G':
			val = 1
		case 'B':
			val = 2
		}
		val = 1 << val
		state[rings[i+1]-'0'] = val | state[rings[i+1]-'0']
	}
	result := 0
	for _, s := range state {
		if s == 7 {
			result++
		}
	}
	return result
}
