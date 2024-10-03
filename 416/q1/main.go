package main

func main() {
}

func reportSpam(message []string, bannedWords []string) bool {
	bannedWordsMap := make(map[string]bool)
	for _, word := range bannedWords {
		bannedWordsMap[word] = true
	}

	cnt := 0
	for _, word := range message {
		if bannedWordsMap[word] {
			cnt++
		}
		if cnt >= 2 {
			return true
		}
	}
	return false
}
