package main

func main() {
}

func maximumLength(s string) int {
	result := -1
	count := 0
	countMap := make(map[rune]map[int]int)
	for idx, ch := range s {
		if idx > 0 && ch != rune(s[idx-1]) {
			count = 0
		}
		if _, ok := countMap[ch]; !ok {
			countMap[ch] = make(map[int]int)
		}
		count++
		if count > 2 && result < count-2 {
			result = count - 2
		}
		// add count -1 and count
		countMap[ch][count]++
		if count-1 > 0 {
			countMap[ch][count-1]++
		}
		if countMap[ch][count] >= 3 && result < count {
			result = count
			continue
		}
		if countMap[ch][count-1] >= 3 && result < count-1 {
			result = count - 1
		}
	}
	return result
}
