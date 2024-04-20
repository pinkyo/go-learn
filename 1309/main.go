package main

import "fmt"

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
}

func freqAlphabets(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && s[i+2] == '#' {
			result += string((s[i]-'0')*10 + s[i+1] - '0' - 1 + 'a')
			i += 2
		} else {
			result += string(s[i] - '0' - 1 + 'a')
		}
	}
	return result
}
