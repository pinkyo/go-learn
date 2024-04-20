package main

func main() {

}

func countSeniors(details []string) int {
	var result int
	for _, item := range details {
		age := item[11:13]
		if age > "60" {
			result++
		}
	}
	return result
}
