package main

func main() {

}

func oddString(words []string) string {
	for i := 0; i < len(words[0])-1; i++ {
		cntMap := make(map[int]int)
		idxMap := make(map[int]int)
		for j := 0; j < len(words); j++ {
			cntMap[int(words[j][i+1]-words[j][i])]++
			idxMap[int(words[j][i+1]-words[j][i])] = j
		}
		for diff, cnt := range cntMap {
			if cnt == 1 {
				return words[idxMap[diff]]
			}
		}
	}
	return ""
}
