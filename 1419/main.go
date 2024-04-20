package main

import "fmt"

func main() {
	fmt.Println(minNumberOfFrogs("croakcroak"))
}

func minNumberOfFrogs(croakOfFrogs string) int {
	preChar := make([]int, 5)
	var result int
	var cur int
	for i := range croakOfFrogs {
		ch := croakOfFrogs[i]
		idx := getIdxOfChar(ch)
		if idx == -1 {
			return -1
		}
		if idx == 0 {
			preChar[idx]++
			cur++
			if cur > result {
				result = cur
			}
		} else {
			preChar[idx-1]--
			preChar[idx]++
			if preChar[idx-1] < 0 {
				return -1
			}
		}
		if idx == 4 {
			cur--
		}
	}

	for i := 0; i < 4; i++ {
		if preChar[i] != 0 {
			return -1
		}
	}
	return result
}

func getIdxOfChar(ch byte) int {
	switch ch {
	case 'c':
		return 0
	case 'r':
		return 1
	case 'o':
		return 2
	case 'a':
		return 3
	case 'k':
		return 4
	}
	return -1
}
