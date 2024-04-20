package main

func main() {

}

func minimizedStringLength(s string) int {
	charMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charMap[s[i]] = charMap[s[i]] + 1
	}
	return len(charMap)
}
