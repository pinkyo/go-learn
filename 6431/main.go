package main

func main() {

}

func doesValidArrayExist(derived []int) bool {
	var result int
	for i := 0; i < len(derived); i++ {
		result ^= derived[i]
	}
	return result == 0
}
