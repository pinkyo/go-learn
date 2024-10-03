package main

func main() {
}

func countKConstraintSubstrings(s string, k int) int {
	result := 0
	for i := 0; i < len(s); i++ {
		count1 := 0
		count0 := 0
		for j := i; j < len(s); j++ {
			if s[j] == '1' {
				count1++
			} else {
				count0++
			}
			if count0 <= k || count1 <= k {
				result++
			} else if count0 > k && count1 > k {
				break
			}
		}
	}
	return result
}
