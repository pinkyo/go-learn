package main

func main() {
}

func findLatestTime(s string) string {
	if s[0] == '?' {
		if s[1] == '?' || s[1] < '2' {
			s = "1" + s[1:]
		} else {
			s = "0" + s[1:]
		}
	}

	if s[1] == '?' {
		if s[0] == '0' {
			s = "09" + s[2:]
		} else {
			s = "11" + s[2:]
		}
	}

	if s[3] == '?' {
		s = s[:3] + "5" + s[4:]
	}

	if s[4] == '?' {
		s = s[:4] + "9"
	}

	return s
}
