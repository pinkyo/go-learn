package main

func main() {
}

func canMakeSubsequence(str1 string, str2 string) bool {
	idx2 := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[idx2] || (str1[i]-'a'+1)%26 == str2[idx2]-'a' {
			idx2++
		}
		if idx2 == len(str2) {
			return true
		}
	}
	return false
}
