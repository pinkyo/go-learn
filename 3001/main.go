package main

import "fmt"

func main() {
	fmt.Println(minMovesToCaptureTheQueen(4, 3, 3, 4, 2, 5))
	// fmt.Println(minMovesToCaptureTheQueen(4, 3, 3, 4, 5, 2))
	// fmt.Println(minMovesToCaptureTheQueen(1, 1, 1, 4, 1, 8))
}

func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e && c != e {
		return 1
	}
	if a == e && c == e && !(b > d && d > f || b < d && d < f) {
		return 1
	}
	if b == f && d != f {
		return 1
	}
	if b == f && d == f && !(a > c && c > e || a < c && c < e) {
		return 1
	}
	xx := f - d
	yy := e - c
	if abs(xx) == abs(yy) && xx*(e-a) != yy*(f-b) {
		return 1
	}
	if abs(xx) == abs(yy) && xx*(e-a) == yy*(f-b) && !(c < a && a < e || e < a && a < c) {
		return 1
	}
	return 2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
