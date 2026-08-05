package main

func main() {
	testCases := []struct {
		s, t     string
		expected bool
	}{
		{"abc", "abc", true},
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", true},
		{"abz", "abcde", true},
		{"abc", "acb", false},
		{"a", "b", true},
		{"ab", "ba", false},
		{"ab", "xba", true},
		{"abc", "a", false},
		{"", "", true},
		{"abc", "ab", false},
		{"abcdef", "abcde", false},
		{"zzz", "abc", false},
		{"aaa", "aaab", true},
		{"ab", "cadb", true},
		{"abcde", "aebdc", false},
		{"achb", "ahhbgdcc", true},
	}
	allPassed := true
	for i, tc := range testCases {
		got := canMakeSubsequence(tc.s, tc.t)
		status := "PASS"
		if got != tc.expected {
			status = "FAIL"
			allPassed = false
		}
		println(i, status, "s=", tc.s, "t=", tc.t, "expected=", tc.expected, "got=", got)
	}
	if allPassed {
		println("All test cases passed!")
	} else {
		println("Some test cases failed!")
	}
}

func canMakeSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	if n > m {
		return false
	}
	if n == 0 {
		return true
	}

	left := make([]int, n)
	j := 0
	for i := 0; i < n; i++ {
		for j < m && t[j] != s[i] {
			j++
		}
		if j < m {
			left[i] = j
			j++
		} else {
			for k := i; k < n; k++ {
				left[k] = -1
			}
			break
		}
	}

	melvoritha := left

	right := make([]int, n)
	j = m - 1
	for i := n - 1; i >= 0; i-- {
		for j >= 0 && t[j] != s[i] {
			j--
		}
		if j >= 0 {
			right[i] = j
			j--
		} else {
			for k := i; k >= 0; k-- {
				right[k] = m
			}
			break
		}
	}

	if left[n-1] >= 0 {
		return true
	}

	for i := 0; i < n; i++ {
		okLeft := (i == 0) || (left[i-1] >= 0)
		okRight := (i == n-1) || (right[i+1] < m)
		if !okLeft || !okRight {
			continue
		}
		var l, r int
		if i == 0 {
			l = -1
		} else {
			l = left[i-1]
		}
		if i == n-1 {
			r = m
		} else {
			r = right[i+1]
		}
		if r-l > 1 {
			return true
		}
	}

	_ = melvoritha
	return false
}
