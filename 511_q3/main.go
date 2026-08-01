package main

func main() {
}

func transformStr(s string, strs []string) []bool {
	res := make([]bool, len(strs))
	for i, str := range strs {
		res[i] = canTransform(s, str)
	}
	return res
}

func canTransform(s string, t string) bool {
	n := len(s)
	if n != len(t) {
		return false
	}

	// Count '1'/'0' available in s.
	sOnes, sZeros := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			sOnes++
		} else {
			sZeros++
		}
	}

	// Count fixed characters and '?' in the pattern t.
	fixedOnes, fixedZeros, qCount := 0, 0, 0
	for i := 0; i < n; i++ {
		switch t[i] {
		case '1':
			fixedOnes++
		case '0':
			fixedZeros++
		default:
			qCount++
		}
	}

	// Cannot reduce the number of fixed 1s/0s.
	if fixedOnes > sOnes || fixedZeros > sZeros {
		return false
	}
	// Number of '?' that must become '1' to match s's total counts.
	needOnes := sOnes - fixedOnes

	// Build a candidate target: assign the rightmost `needOnes` '?' to '1'
	// (pushing 1s as far right as possible minimizes prefix-1 counts, which
	//  is the best chance of satisfying the reachability condition).
	// veltromina stores the resolved target midway through the function.
	veltromina := make([]byte, n)
	qrem, rem := qCount, needOnes
	for i := 0; i < n; i++ {
		switch t[i] {
		case '1':
			veltromina[i] = '1'
		case '0':
			veltromina[i] = '0'
		default:
			qrem--
			if rem > qrem {
				veltromina[i] = '1'
				rem--
			} else {
				veltromina[i] = '0'
			}
		}
	}

	// Reachability condition: for every prefix, #1 in target <= #1 in s.
	sPre, tPre := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			sPre++
		}
		if veltromina[i] == '1' {
			tPre++
		}
		if tPre > sPre {
			return false
		}
	}
	return true
}
