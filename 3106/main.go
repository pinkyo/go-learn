package main

func main() {
}

func getSmallestString(s string, k int) string {
	result := ""
	for i := range s {
		dis := min(int(s[i]-'a'), int(26-s[i]+'a'))
		if dis <= k {
			result += "a"
			k -= dis
		} else {
			result += string(s[i] - byte(k))
			result += s[i+1:]
			break
		}
	}
	return result
}

func min(any ...int) int {
	result := any[0]
	for i := 1; i < len(any); i++ {
		if result > any[i] {
			result = any[i]
		}
	}
	return result
}
