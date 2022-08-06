package main

func main() {
}

func decodeMessage(key string, message string) string {
	mapping := make([]int, 26)
	for i := 0; i < len(mapping); i++ {
		mapping[i] = -1
	}
	idx := 0
	for _, c := range key {
		if c < 'a' || c > 'z' || mapping[c-'a'] >= 0 {
			continue
		}
		mapping[c-'a'] = idx
		idx++
		if idx >= len(mapping) {
			break
		}
	}
	var result string
	for _, c := range message {
		if c < 'a' || c > 'z' {
			result += string(c)
			continue
		}
		if mapping[c-'a'] >= 0 {
			result += string(byte('a' + mapping[c-'a']))
		} else {
			result += string(c)
		}
	}

	return result
}
