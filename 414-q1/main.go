package main

import (
	"strconv"
	"strings"
)

func main() {
}

func convertDateToBinary(date string) string {
	fields := strings.Split(date, "-")
	for i, val := range fields {
		val, _ := strconv.Atoi(val)
		fields[i] = strconv.FormatInt(int64(val), 2)
	}

	return strings.Join(fields, "-")
}
