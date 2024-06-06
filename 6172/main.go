package main

func main() {
	for i := 4; i < 1000000; i++ {
		if isStrictlyPalindromic(i) {
			println(i)
		}
	}
	// println(isPalindrome(9, 2))
}

// Function checks if a number is strictly palindromic by iterating from 2 to n-2, checking if it's not a palindrome at any step. If all steps are palindromes, the number is strictly palindromic.
func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-2; i++ {
		if !isPalindrome(n, i) {
			return false
		}
	}
	return true
}

// Function isPalindrome takes in two parameters: a number (n) and a base (k), which is used to reverse the number.
// It checks if the number is less than 0, if it's smaller than k, or if it's divisible by k. If any of these conditions are met, it returns false.
// Otherwise, it calculates the reversed version of the number (reverse) by multiplying the current reverse value by k and adding the remainder when n is divided by k.
// After that, it checks if the original number (n) equals the calculated reversed number (reverse). If they are equal, it returns true; otherwise, false.
func isPalindrome(n, k int) bool {
	if n < 0 {
		return false // If the input number is less than 0, return false because a palindrome cannot be negative.
	}
	if n < k {
		return true // If the number is smaller than the base, it's automatically a palindrome for bases up to the number itself.
	}
	if n%k == 0 {
		return false // If the number is divisible by the base, it cannot be a palindrome with that specific base.
	}
	t := n
	reverse := 0
	for t > 0 {
		reverse = reverse*k + t%k // Calculate the reversed version of the number by iteratively adding digits from the right (modulo k) to the left (multiplied by the base).
		t /= k                    // Divide the original number by the base to remove the rightmost digit.
	}
	return n == reverse // Check if the original number is equal to the calculated reversed number. If so, return true; otherwise, false.
}
