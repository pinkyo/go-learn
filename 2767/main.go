package main

import "fmt"

func main() {
	fmt.Println(minimumBeautifulSubstrings("1011"))
}

func minimumBeautifulSubstrings(s string) int {
	values := make(map[int]struct{})
	value := 1
	values[value] = struct{}{}
	for value < 1<<16-1 {
		value *= 5
		values[value] = struct{}{}
	}
	// fmt.Println("values:", values)
	cache := make(map[string]int)
	var helper func(string) int
	helper = func(str string) int {
		if str == "" {
			return 0
		}
		if str[0] == '0' {
			return -1
		}
		if cached, ok := cache[str]; ok {
			return cached
		}
		result := -1
		current := 0
		for i := 0; i < len(str); i++ {
			current *= 2
			if str[i] == '1' {
				current++
			}
			if _, ok := values[current]; ok {
				if res := helper(str[i+1:]); res >= 0 {
					if result == -1 || result > res+1 {
						result = res + 1
					}
				}
			}
		}
		cache[str] = result
		return result
	}
	// fmt.Println("cache:", cache)
	return helper(s)
}
