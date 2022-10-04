package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countDaysTogether("08-15",
		"08-18",
		"08-16",
		"08-19"))
}

var _days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	beginDate := arriveAlice
	if arriveBob > beginDate {
		beginDate = arriveBob
	}
	endDate := leaveAlice
	if leaveBob < endDate {
		endDate = leaveBob
	}
	if beginDate > endDate {
		return 0
	}
	return calDays(beginDate, endDate)
}

func calDays(begin, to string) int {
	var result int
	beginArr := strings.Split(begin, "-")
	endArr := strings.Split(to, "-")
	beginMonth, _ := strconv.Atoi(beginArr[0])
	beginDay, _ := strconv.Atoi(beginArr[1])
	endMonth, _ := strconv.Atoi(endArr[0])
	endDay, _ := strconv.Atoi(endArr[1])
	if beginMonth == endMonth {
		return endDay - beginDay + 1
	}
	result += endDay
	result += _days[beginMonth-1] - beginDay + 1
	for i := beginMonth + 1; i < endMonth; i++ {
		result += _days[i-1]
	}
	return result
}
