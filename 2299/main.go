package main

import "strings"

func main() {

}

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false
	for i := 0; i < len(password); i++ {
		ch := password[i]
		if i > 0 && ch == password[i-1] {
			return false
		}
		if ch >= 'a' && ch <= 'z' {
			hasLower = true
		} else if ch >= 'A' && ch <= 'Z' {
			hasUpper = true
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		} else if strings.IndexByte("!@#$%^&*()-+", ch) >= 0 {
			hasSpecial = true
		}
	}
	return hasLower && hasUpper && hasDigit && hasSpecial
}
