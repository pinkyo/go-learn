package main

import (
	"strconv"
	"strings"
)

func main() {
}

func secondsBetweenTimes(startTime string, endTime string) int {
	startTimeFields := strings.Split(startTime, ":")
	endTimeFields := strings.Split(endTime, ":")

	hs, _ := strconv.Atoi(startTimeFields[0])
	he, _ := strconv.Atoi(endTimeFields[0])
	ms, _ := strconv.Atoi(startTimeFields[1])
	me, _ := strconv.Atoi(endTimeFields[1])
	ss, _ := strconv.Atoi(startTimeFields[2])
	se, _ := strconv.Atoi(endTimeFields[2])
	result := he - hs
	if result < 0 {
		result += 24
	}
	result *= 60
	result += me - ms
	result *= 60
	result += se - ss
	return result
}
