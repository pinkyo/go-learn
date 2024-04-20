package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	var str string
	for i := 0; i < 50; i++ {
		str += string('a' + i%26)
	}
	fmt.Println("str:", str)
	var dict []string
	var cur int
	for cur < len(str) {
		// generate radom int
		l := rand.Intn(10)
		if cur+l > len(str) {
			l = len(str) - cur
		}
		dict = append(dict, str[cur:cur+l])
		cur += l
	}
	fmt.Println("dict:", dict)
	fmt.Println(minExtraChar(str+"aa", dict))
	// fmt.Println(minExtraChar("leetleet", []string{"leet"}))
	// fmt.Println(minExtraChar("leetscode", []string{"leet", "code", "leetcode"}))
	// fmt.Println(minExtraChar("sayhelloworld", []string{"hello", "world"}))
	// fmt.Println(minExtraChar("dwmodizxvvbosxxw", []string{"ox", "lb", "diz", "gu", "v", "ksv", "o", "nuq", "r", "txhe", "e", "wmo", "cehy", "tskz", "ds", "kzbu"}))
	// fmt.Println(minExtraChar("voctvochpgutoywpnafylzelqsnzsbandjcqdciyoefi",
	// 	[]string{"tf", "v", "wadrya", "a", "cqdci", "uqfg", "voc", "zelqsn", "band", "b", "yoefi", "utoywp", "herqqn", "umra", "frfuyj", "vczatj", "sdww"}))
}

func minExtraChar(s string, dictionary []string) int {
	sort.Slice(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})
	var min int = 1<<31 - 1
	var helper func(strs []string, dicts []string, pos int)
	helper = func(strs []string, dicts []string, pos int) {
		if pos == len(dicts) {
			var result int
			for _, str := range strs {
				result += len(str)
			}
			if result < min {
				fmt.Println("strs:", strs, "pos:", pos, "len:", len(strs), "result:", result)
				min = result
			}
			return
		}
		word := dicts[pos]
		for idx, str := range strs {
			p := 0
			pp := strings.Index(str[p:], word)
			for pp >= 0 {
				var newStr []string
				newStr = append(newStr, strs[:idx]...)
				if str[:p+pp] != "" {
					newStr = append(newStr, str[:p+pp])
				}
				if str[p+pp+len(word):] != "" {
					newStr = append(newStr, str[p+pp+len(word):])
				}
				newStr = append(newStr, strs[idx+1:]...)
				helper(newStr, dicts, pos+1)
				p += pp + 1
				if p >= len(str) {
					break
				}
				pp = strings.Index(str[p:], word)
			}
		}
		helper(strs, dicts, pos+1)
	}
	helper([]string{s}, dictionary, 0)
	return min
}
