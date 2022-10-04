package main

func main() {
	for i := 4; i < 1000000; i++ {
		if isStrictlyPalindromic(i) {
			println(i)
		}
	}
	// println(isPalindrome(9, 2))
}

func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-2; i++ {
		if !isPalindrome(n, i) {
			return false
		}
	}
	return true
}

func isPalindrome(n, k int) bool {
	if n < 0 {
		return false
	}
	if n < k {
		return true
	}
	if n%k == 0 {
		return false
	}
	t := n
	reverse := 0
	for t > 0 {
		reverse = reverse*k + t%k
		t /= k
	}
	return n == reverse
}
