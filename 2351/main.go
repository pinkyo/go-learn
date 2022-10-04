package main

func main() {
}

func repeatedCharacter(s string) byte {
	cntMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		cntMap[s[i]]++
		if cntMap[s[i]] == 2 {
			return s[i]
		}
	}
	return 0
}
