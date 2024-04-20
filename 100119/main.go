package main

import "fmt"

func main() {
	// fmt.Println(maximumXorProduct(12, 5, 4))
	fmt.Println(maximumXorProduct(1, 6, 3))
}

func maximumXorProduct(a int64, b int64, n int) int {
	mod := int64(1e9 + 7)
	for i := n - 1; i >= 0; i-- {
		mask := int64(1 << i)
		if a^mask == 0 || b^mask == 0 || a&mask == 1 && b&mask == 1 {
			continue
		}
		if a == 0 || b == 0 || a&mask == 0 && b&mask == 0 || float64(a)/float64(a^mask) < float64(b^mask)/float64(b) {
			a ^= mask
			b ^= mask
		}
	}
	a %= mod
	b %= mod
	return int(a * b % mod)
}
