package main

func main() {
}

func isPossibleToRearrange(s string, t string, k int) bool {
	cntMap := make(map[string]int)
	l := len(s) / k
	for i := 0; i < len(s); i += l {
		cntMap[s[i:i+l]]++
	}
	for i := 0; i < len(t); i += l {
		cntMap[t[i:i+l]]--
	}
	for _, cnt := range cntMap {
		if cnt < 0 {
			return false
		}
	}
	return true
}
