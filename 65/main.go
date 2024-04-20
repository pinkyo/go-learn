package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// fmt.Println(isNumber("-90E3"))
	fmt.Println(isNumber("..2"))
}

func isNumber(s string) bool {
	// 处理E/e大小写的问题
	lowerS := strings.ToLower(s)
	fields := strings.Split(lowerS, "e")
	if len(fields) > 2 {
		// 多个e的情况
		return false
	}
	if !isInteger(fields[0]) && !isFloat(fields[0]) {
		return false
	}
	if len(fields) == 2 && !isInteger(fields[1]) {
		// 判断e前面和后面是否是数字
		return false
	}
	return true
}

// 判断是否是整数
func isInteger(s string) bool {
	pattern := "^[-\\+]?[0-9]+$"
	matched, _ := regexp.Match(pattern, []byte(s))
	return matched
}

// 判断是否是小数
func isFloat(s string) bool {
	pattern := "^[-\\+]?(([0-9]+\\.[0-9]*)|([0-9]*\\.[0-9]+))$"
	matched, _ := regexp.Match(pattern, []byte(s))
	return matched
}
