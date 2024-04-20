package main

func main() {

}

func minimumSteps(s string) int64 {
	pivot := 0
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			continue
		}
		result += (i - pivot)
		pivot++
	}
	return int64(result)
}
