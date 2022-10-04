package main

func main() {
}

func maxLengthBetweenEqualCharacters(s string) int {
	result := -1
	posMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if pos, ok := posMap[s[i]]; ok {
			result = max(result, i-pos-1)
		} else {
			posMap[s[i]] = i
		}
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
