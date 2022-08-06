package main

import "fmt"

func main() {
	fmt.Println(matchReplacement("fool3e7bar", "leet",
		[][]byte{{'e', '3'}, {'t', '7'}, {'t', '8'}}))
}

func matchReplacement(s string, sub string, mappings [][]byte) bool {
	chMap := make(map[byte]map[byte]struct{})
	for _, v := range mappings {
		if _, ok := chMap[v[0]]; !ok {
			chMap[v[0]] = make(map[byte]struct{})
		}
		chMap[v[0]][v[1]] = struct{}{}
	}
OUTER:
	for i := 0; i < len(s)-len(sub)+1; i++ {
		for j := 0; j < len(sub); j++ {
			if s[i+j] == sub[j] {
				continue
			}
			if _, ok := chMap[sub[j]][s[i+j]]; ok {
				continue
			}
			continue OUTER
		}
		return true
	}
	return false
}
