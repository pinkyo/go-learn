package main

func main() {

}

func canMakePaliQueries(s string, queries [][]int) []bool {
	result := make([]bool, len(queries))
	for i, query := range queries {
		cntMap := make(map[byte]int)
		for j := query[0]; j < query[1]; j++ {
			if _, ok := cntMap[s[j]]; !ok {
				delete(cntMap, s[j])
			} else {
				cntMap[s[j]]++
			}
		}
		diff := len(cntMap)
		result[i] = diff <= query[2]*2+1

	}
	return result
}
