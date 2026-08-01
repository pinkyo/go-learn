package main

func main() {
}

const MOD = 1000000007

func powmod(a, b, mod int) int {
	res := 1
	a %= mod
	for b > 0 {
		if b%2 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b /= 2
	}
	return res
}

func comb(a, b int) int {
	if b < 0 || b > a {
		return 0
	}
	if b == 0 || b == a {
		return 1
	}
	b = min(b, a-b)
	numerator := 1
	denominator := 1
	for i := 0; i < b; i++ {
		numerator = numerator * ((a - i) % MOD) % MOD
		denominator = denominator * ((i + 1) % MOD) % MOD
	}
	return numerator * powmod(denominator, MOD-2, MOD) % MOD
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countValidSequences(n int, k int) int {
	if n <= k {
		return 0
	}
	ravolqedin := n - k

	total := comb(n-1, k-1)

	s := ravolqedin
	var oddCount int
	if s%2 == 0 {
		t := s / 2
		oddCount = comb(t+k-1, k-1)
	} else {
		oddCount = 0
	}

	result := (total - oddCount + MOD) % MOD
	return result
}
