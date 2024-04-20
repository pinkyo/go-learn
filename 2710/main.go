package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeTrailingZeros("123000"))
}

func removeTrailingZeros(num string) string {
	return strings.TrimRight(num, "0")
}
