package main

func main() {
}

func minOperations(s string) int {
	cur := []int{0, 1}
	var cnt [2]int
	for i := 0; i < len(s); i++ {
		if int(s[i]-'0') != cur[0] {
			cnt[0]++
		}
		if int(s[i]-'0') != cur[1] {
			cnt[1]++
		}
		cur[0] = (cur[0] + 1) % 2
		cur[1] = (cur[1] + 1) % 2
	}
	if cnt[0] > cnt[1] {
		return cnt[1]
	}
	return cnt[0]
}
