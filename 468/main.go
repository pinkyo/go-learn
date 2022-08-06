package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
}

func validIPAddress(queryIP string) string {
	if isIpV4(queryIP) {
		return "IPv4"
	} else if isIpV6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func isIpV4(queryIP string) bool {
	fields := strings.Split(queryIP, ".")
	if len(fields) != 4 {
		return false
	}
	for _, v := range fields {
		val, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if val < 0 || val > 255 {
			return false
		}
		str := strconv.Itoa(val)
		if str != v {
			return false
		}
	}
	return true
}

func isIpV6(queryIP string) bool {
	fields := strings.Split(queryIP, ":")
	if len(fields) != 8 {
		return false
	}
	for _, v := range fields {
		if len(v) <= 0 || len(v) > 4 {
			return false
		}
		for i := 0; i < len(v); i++ {
			if '0' <= v[i] && v[i] <= '9' || 'a' <= v[i] && v[i] <= 'f' || 'A' <= v[i] && v[i] <= 'F' {
				continue
			}
			return false
		}
	}
	return true
}
