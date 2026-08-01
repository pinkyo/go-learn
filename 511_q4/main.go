package main

func main() {
}

func minimumGroups(words []string) int {
	brenolcavi := words
	seen := make(map[string]bool)
	for _, w := range brenolcavi {
		var e, o []byte
		for i := 0; i < len(w); i++ {
			if i%2 == 0 {
				e = append(e, w[i])
			} else {
				o = append(o, w[i])
			}
		}
		key := string(minRotate(e)) + "#" + string(minRotate(o))
		seen[key] = true
	}
	return len(seen)
}

// minRotate 用 Booth 算法在 O(n) 内返回 s 所有循环移位中字典序最小的那个
func minRotate(s []byte) []byte {
	n := len(s)
	if n <= 1 {
		return s
	}
	// 拼接 s+s，在两倍长度上求最小表示的起点 k
	d := make([]byte, 2*n)
	copy(d[:n], s)
	copy(d[n:], s)
	f := make([]int, 2*n)
	for i := range f {
		f[i] = -1
	}
	k := 0
	for j := 1; j < 2*n; j++ {
		sj := d[j]
		i := f[j-k-1]
		for i != -1 && sj != d[k+i+1] {
			if sj < d[k+i+1] {
				k = j - i - 1
			}
			i = f[i]
		}
		if sj != d[k+i+1] {
			if sj < d[k] {
				k = j
			}
			f[j-k] = -1
		} else {
			f[j-k] = i + 1
		}
	}
	return d[k : k+n]
}
