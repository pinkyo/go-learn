package main

import "strconv"

func main() {
}

func countTime(time string) int {
	var h, m int
	switch {
	case time[0] != '?' && time[1] != '?':
		hour, _ := strconv.Atoi(time[:2])
		if hour > 23 {
			return 0
		}
		h = 1
	case time[0] != '?' && time[1] == '?':
		hour, _ := strconv.Atoi(time[:1])
		if hour > 2 {
			return 0
		}
		if hour == 2 {
			h = 4
		} else {
			h = 10
		}
	case time[0] == '?' && time[1] != '?':
		hour, _ := strconv.Atoi(time[1:2])
		if hour > 3 {
			h = 2
		} else {
			h = 3
		}
	case time[0] == '?' && time[1] == '?':
		h = 24
	}

	switch {
	case time[3] != '?' && time[4] != '?':
		minute, _ := strconv.Atoi(time[3:5])
		if minute > 59 {
			return 0
		}
		m = 1
	case time[3] != '?' && time[4] == '?':
		minute, _ := strconv.Atoi(time[3:4])
		if minute > 5 {
			return 0
		}
		m = 10
	case time[3] == '?' && time[4] != '?':
		minute, _ := strconv.Atoi(time[4:5])
		if minute > 9 {
			return 0
		}
		m = 6
	case time[3] == '?' && time[4] == '?':
		m = 60
	}
	return h * m
}
