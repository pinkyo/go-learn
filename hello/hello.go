package main

import (
	"archive/tar"
	"archive/zip"
	"container/heap"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

func main() {
	defer recordTime()()
	// arr := make([]int, 8)
	// for i := 1; i <= 8; i++ {
	// 	func() {
	// 		defer recordTime()()
	// 		res := largestPalindrome(i)
	// 		arr[i-1] = res
	// 		fmt.Printf("i=%d, res=%d\n", i, res)
	// 	}()
	// }
	// fmt.Println(arr)
	// bs := []byte("abcdefghijklmnopqrstuvwxyz")
	// bytesToString: time used: 0.222800
	// bytesToString2: time used: 2.911676
	// for i := 0; i < 1; i++ {
	// 	bytesToString2(bs)
	// }
	// fmt.Println(checkValid([][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}))
	// fmt.Println(math.MaxInt64)
	// var intVal int = math.MaxInt64
	// fmt.Println(intVal)
	// fmt.Println(unsafe.Sizeof(intVal))
	// fmt.Println(longestPath([]int{-1, 0, 0, 1, 1, 2}, "abacbe"))

	ch := make(chan string, 1)
	ch <- "abc"
	close(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(1)
}

func findDuplicates(nums []int) []int {
	var result []int
	eixst := make([]bool, 100001)
	for i := 0; i < len(nums); i++ {
		if eixst[nums[i]] {
			result = append(result, nums[i])
		} else {
			eixst[nums[i]] = true
		}
	}
	return result
}

func longestPath(parent []int, s string) int {
	var result int
	arr := make([]int, len(parent))
	childCnt := make(map[int]int, len(parent))
	for i := 0; i < len(parent); i++ {
		if parent[i] != -1 {
			childCnt[parent[i]]++
		}
	}
	var nodes []int
	for i := 0; i < len(parent); i++ {
		if _, ok := childCnt[i]; !ok {
			nodes = append(nodes, i)
		}
	}
	for len(nodes) > 0 {
		var next []int
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			if parent[node] == -1 {
				break
			}
			if s[node] != s[parent[node]] {
				val := arr[node] + 1 + arr[parent[node]]
				if val > result {
					result = val
				}
				if arr[node]+1 > arr[parent[node]] {
					arr[parent[node]] = arr[node] + 1
				}
			}
			childCnt[parent[node]]--
			if childCnt[parent[node]] == 0 {
				next = append(next, parent[node])
			}
		}
		nodes = next
	}
	return result + 1
}

func maxTrailingZeros(grid [][]int) int {
	var result int
	cal := func(ii, jj int, fi, fj func(int) int) {
		cntArr10 := make([]int, len(grid[0]))
		cntArr5 := make([]int, len(grid[0]))
		cntArr2 := make([]int, len(grid[0]))
		for i := ii; 0 <= i && i < len(grid); i = fi(i) {
			var cnt10 int
			var cnt5 int
			var cnt2 int
			for j := jj; 0 <= j && j < len(grid[0]); j = fj(j) {
				var c10 int
				var c5 int
				var c2 int
				val := grid[i][j]
				for val > 0 && val%10 == 0 {
					c10++
					val /= 10
				}
				for val > 0 && val%5 == 0 {
					c5++
					val /= 5
				}
				for val > 0 && val%2 == 0 {
					c2++
					val /= 2
				}
				cntArr10[j] += c10
				cntArr5[j] += c5
				cntArr2[j] += c2
				cc10 := cnt10 + cntArr10[j]
				cc5 := cnt5 + cntArr5[j]
				cc2 := cnt2 + cntArr2[j]
				val = cc10 + min(cc5, cc2)
				if val > result {
					result = val
				}
				cnt10 += c10
				cnt5 += c5
				cnt2 += c2
			}
		}
	}
	cal(0, 0, func(i int) int { return i + 1 }, func(j int) int { return j + 1 })
	cal(len(grid)-1, 0, func(i int) int { return i - 1 }, func(j int) int { return j + 1 })
	cal(0, len(grid[0])-1, func(i int) int { return i + 1 }, func(j int) int { return j - 1 })
	cal(len(grid)-1, len(grid[0])-1, func(i int) int { return i - 1 }, func(j int) int { return j - 1 })
	return result
}

func minimumRounds(tasks []int) int {
	cntMap := make(map[int]int)
	for _, v := range tasks {
		cntMap[v]++
	}
	var result int
	for _, v := range cntMap {
		if v < 2 {
			return -1
		}
		result += v / 3
		if v%3 > 0 {
			result++
		}
	}
	return result
}

func digitSum(s string, k int) string {
	cur := s
	for len(cur) > k {
		var sum int
		var str string
		for i := 0; i < len(cur); i++ {
			if i > 0 && i%k == 0 {
				str += strconv.Itoa(int(sum))
				sum = 0
			}
			sum += int(cur[i] - '0')
		}
		str += strconv.Itoa(int(sum))
		cur = str
	}
	return cur
}

func appealSum(s string) int64 {
	result := 0
	preMap := make(map[byte]int, 0)
	cur := 0
	for i := 0; i < len(s); i++ {
		if idx, ok := preMap[s[i]]; !ok {
			cur += i + 1
		} else {
			cur += i - idx
		}
		result += cur
		preMap[s[i]] = i
	}
	return int64(result)
}

func countDistinct(nums []int, k int, p int) int {
	var result int
	existMap := make(map[string]struct{})
	dividable := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dividable[i] = nums[i]%p == 0
	}
	for i := 0; i < len(nums); i++ {
		var str string
		var sep string
		var c = 0
		for j := i + 1; j < len(nums); j++ {
			if dividable[j] {
				c++
			}
			if c > k {
				break
			}
			str = str + sep + strconv.Itoa(nums[j])
			sep = ","
			if _, ok := existMap[str]; !ok {
				result++
				existMap[str] = struct{}{}
			}
		}
	}
	return result
}

func minimumCardPickup(cards []int) int {
	result := len(cards) + 1
	posMap := make(map[int]int)
	for i := 0; i < len(cards); i++ {
		if pre, ok := posMap[cards[i]]; ok {
			result = min(result, i-pre+1)
		}
		posMap[cards[i]] = i
	}
	if result > len(cards) {
		return -1
	}
	return result
}

// func maximumMinutes(grid [][]int) int {
// 	minutes := make([]int, len(grid))
// 	for i := 0; i < len(minutes); i++ {
// 		minutes[i] = len(grid[i])
// 	}
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[i]); j++ {
// 			if grid[i][j] == 1 {

// 			}
// 		}
// 	}
// }

func findTheWinner(n int, k int) int {
	players := make([]int, n)
	for i := 0; i < n; i++ {
		players[i] = (i + 1) % n
	}
	pre := n - 1
	cur := 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < k-1; j++ {
			pre = cur
			cur = players[cur]
		}
		players[pre] = players[cur]
		players[cur] = -1
		cur = players[pre]
	}
	return cur + 1
}

func removeDigit(number string, digit byte) string {
	var result string
	for i := 0; i < len(number); i++ {
		if number[i] != digit {
			continue
		}
		val := number[:i] + number[i+1:]
		if val > result {
			result = val
		}
	}
	return result
}

func reorderLogFiles(logs []string) []string {
	var digit []string
	var letter []string
	for _, v := range logs {
		idx := strings.Index(v, " ")
		if unicode.IsDigit(rune(v[idx+1])) {
			digit = append(digit, v)
		} else {
			letter = append(letter, v)
		}
	}
	sort.Slice(letter, func(i, j int) bool {
		idx1 := strings.Index(letter[i], " ")
		idx2 := strings.Index(letter[j], " ")
		if letter[i][idx1+1:] == letter[j][idx2+1:] {
			return letter[i][:idx1] < letter[j][:idx2]
		}
		return letter[i][idx1+1:] < letter[j][idx2+1:]
	})
	for _, v := range digit {
		letter = append(letter, v)
	}
	return letter
}

func countPrefixes(words []string, s string) int {
	var result int
	for _, v := range words {
		if strings.HasPrefix(s, v) {
			result++
		}
	}
	return result
}

func minimumAverageDifference(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	var cur int
	var min = math.MaxInt32
	var idx = -1
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		val := cur / (i + 1)
		if len(nums)-i-1 > 0 {
			val -= (sum - cur) / (len(nums) - i - 1)
		}
		if val < 0 {
			val = -val
		}
		if val < min {
			min = val
			idx = i
		}
	}
	return idx
}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	for _, v := range walls {
		grid[v[0]][v[1]] = 2
	}
	for _, v := range guards {
		grid[v[0]][v[1]] = 1
	}
	for _, v := range guards {
		idx := v[0] - 1
		for idx >= 0 && grid[idx][v[1]] <= 0 {
			grid[idx][v[1]] = -1
			idx--
		}
		idx = v[0] + 1
		for idx < m && grid[idx][v[1]] <= 0 {
			grid[idx][v[1]] = -1
			idx++
		}
		idx = v[1] - 1
		for idx >= 0 && grid[v[0]][idx] <= 0 {
			grid[v[0]][idx] = -1
			idx--
		}
		idx = v[1] + 1
		for idx < n && grid[v[0]][idx] <= 0 {
			grid[v[0]][idx] = -1
			idx++
		}
	}
	var result int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				result++
			}
		}
	}
	return result
}

func projectionArea(grid [][]int) int {
	var result int
	for _, v := range grid {
		for _, vv := range v {
			if vv > 0 {
				result++
			}
		}
	}
	for _, v := range grid {
		var maxVal int
		for _, vv := range v {
			maxVal = max(maxVal, vv)
		}
		result += maxVal
	}

	for i := 0; i < len(grid[0]); i++ {
		var maxVal int
		for _, v := range grid {
			maxVal = max(maxVal, v[i])
		}
		result += maxVal
	}
	return result
}

func countEven(num int) int {
	var result int
	for num > 0 {
		var sum int
		tmp := num
		for tmp > 0 {
			sum += tmp % 10
			tmp /= 10
		}
		if sum%2 == 0 {
			result++
		}
		num--
	}
	return result
}

func checkValid(matrix [][]int) bool {
	// judge if row of the matrix is valid
	for _, v := range matrix {
		has := make([]bool, len(v))
		cnt := 0
		for _, v2 := range v {
			if 0 <= v2-1 && v2-1 < len(has) && !has[v2-1] {
				has[v2-1] = true
				cnt++
			}
		}
		if cnt != len(has) {
			return false
		}
	}
	// judge if column of the matrix is valid
	for i := 0; i < len(matrix[0]); i++ {
		has := make([]bool, len(matrix[0]))
		cnt := 0
		for _, v := range matrix {
			if 0 <= v[i]-1 && v[i]-1 < len(has) && !has[v[i]-1] {
				has[v[i]-1] = true
				cnt++
			}
		}
		if cnt != len(has) {
			return false
		}
	}
	return true
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	var result int
	for i := 0; i < len(seats); i++ {
		result += abs(seats[i] - students[i])
	}
	return result
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func lengthLongestPath(input string) int {
	maxLen := 0
	stack := []int{0}
	curLen := 0
	for _, v := range strings.Split(input, "\n") {
		level := strings.Count(v, "\t")
		for len(stack) > level {
			curLen -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, len(v)-level)
		curLen += stack[len(stack)-1]
		if strings.Contains(v, ".") {
			maxLen = max(maxLen, curLen+level)
		}
	}
	return maxLen
}

func maxWidthRamp(nums []int) int {
	mArr := make([]int, len(nums))
	mArr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		mArr[i] = min(mArr[i-1], nums[i])
	}

	binarySearch := func(start, end, target int) int {
		for start <= end {
			mid := start + (end-start)/2
			if mArr[mid] > target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
		return start
	}
	result := 0
	for i := len(nums) - 1; i >= 0; i-- {
		res := binarySearch(0, i-1, nums[i])
		if res < i {
			result = max(result, i-res)
		}
	}
	return result
}

// get min
func min(a int, b ...int) int {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}

// get max value
func max(a int, b ...int) int {
	for _, v := range b {
		if v > a {
			a = v
		}
	}
	return a
}

func shortestToChar(s string, c byte) []int {
	result := make([]int, len(s))
	val := len(s) + 1
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			val = 0
		}
		result[i] = val
		val++
	}
	val = len(s) + 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			val = 0
		}
		if result[i] > val {
			result[i] = val
		}
		val++
	}
	return result
}
func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(drinks)
	var result int
	mod := int(1e9 + 7)
	binarySearch := func(start, end, target int) int {
		for start <= end {
			mid := (start + end) / 2
			if drinks[mid] <= target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
		return end
	}
	for i := 0; i < len(staple); i++ {
		res := binarySearch(0, len(drinks)-1, x-staple[i])
		if 0 <= res && res < len(drinks) {
			result += res + 1
		}
		if result > mod {
			result %= mod
		}
	}
	return result
}

func diagonalSort(mat [][]int) [][]int {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			val := mat[i][j]
			k := 1
			for ; k <= i && k <= j; k++ {
				if mat[i-k][j-k] > val {
					mat[i-k+1][j-k+1] = mat[i-k][j-k]
				} else {
					mat[i-k+1][j-k+1] = val
					k = -1
					break
				}
			}
			if k != -1 {
				mat[i-k+1][j-k+1] = val
			}
		}
	}
	return mat
}

func arrangeWords(text string) string {
	strs := strings.Fields(text)
	strs[0] = strings.ToLower(strs[0])
	// stable sort
	sort.SliceStable(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	result := strings.Join(strs, " ")
	return strings.ToUpper(result[0:1]) + result[1:]
}

func mostCommonWord(paragraph string, banned []string) string {
	strs := strings.FieldsFunc(paragraph, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	m := make(map[string]int)
	for _, v := range strs {
		v = strings.ToLower(v)
		m[v]++
	}
	for _, v := range banned {
		delete(m, v)
	}
	var max int
	var res string
	for k, v := range m {
		if v > max {
			max = v
			res = k
		}
	}
	return res
}

func subsetXORSum(nums []int) int {
	limit := 1 << len(nums)
	var result int
	for i := 0; i < limit; i++ {
		var val int
		state := i
		for i := 0; i < 12 && state != 0; i++ {
			if state&1 == 1 {
				val ^= nums[i]
			}
			state >>= 1
		}
		result += val
	}
	return result
}

func maximumScore(scores []int, edges [][]int) int {
	tree := make(map[int]map[int]bool)
	for _, v := range edges {
		if _, ok := tree[v[0]]; !ok {
			tree[v[0]] = make(map[int]bool)
		}
		tree[v[0]][v[1]] = true
		if _, ok := tree[v[1]]; !ok {
			tree[v[1]] = make(map[int]bool)
		}
		tree[v[1]][v[0]] = true
	}
	var dfs func(int, int, int) int
	visited := make([]bool, len(scores))
	// cache := make(map[int]int)
	dfs = func(pre, node, depth int) int {
		if depth == 0 {
			return scores[node]
		}
		// state := (pre << 33) + (node << 2) + depth
		// if _, ok := cache[state]; ok {
		// 	return cache[state]
		// }
		visited[node] = true
		max := -1
		for child := range tree[node] {
			if visited[child] {
				continue
			}
			res := dfs(node, child, depth-1)
			if res > max {
				max = res
			}
		}
		visited[node] = false
		if max > 0 {
			max += scores[node]
		}
		// cache[state] = max
		return max
	}
	result := -1
	for i := 0; i < len(scores); i++ {
		res := dfs(i, i, 3)
		if res > result {
			result = res
		}
	}
	return result
}

func recordTime() func() {
	startTime := time.Now()
	return func() {
		fmt.Printf("time used: %f\n", time.Since(startTime).Seconds())
	}
}

func findClosestNumber(nums []int) int {
	result := math.MaxInt32
	for _, v := range nums {
		if math.Abs(float64(result)) > math.Abs(float64(v)) ||
			math.Abs(float64(result)) == math.Abs(float64(v)) && result < v {
			result = v
		}
	}
	return result
}

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	result := 0
	for i := 0; i <= total/cost1; i++ {
		result += (total-cost1*i)/cost2 + 1
	}
	return int64(result)
}

type ATM struct {
	count []int
}

var amountArr []int = []int{20, 50, 100, 200, 500}

func Constructor() ATM {
	return ATM{count: make([]int, 5)}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < len(banknotesCount); i++ {
		this.count[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	result := make([]int, 5)
	for i := len(amountArr) - 1; i >= 0; i-- {
		if amount >= amountArr[i] {
			result[i] = amount / amountArr[i]
		}
		if result[i] > this.count[i] {
			result[i] = this.count[i]
		}
		amount -= result[i] * amountArr[i]
	}
	if amount > 0 {
		return []int{-1}
	}
	for i := 0; i < len(result); i++ {
		this.count[i] -= result[i]
	}
	return result
}

const (
	a = 1 << 70
	f = 1e10000
)

type Test struct {
	I int
	F float32
}

func (t *Test) String() string {
	return "test"
}

func test() {
	fmt.Println("test")
}

func missingRolls(rolls []int, mean int, n int) []int {
	sum := 0
	for _, v := range rolls {
		sum += v
	}
	r := mean*(len(rolls)+n) - sum
	rr := r % n
	val := r / n
	if r < n || val > 6 || (val == 6 && rr > 0) {
		return []int{}
	}
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i < rr {
			result = append(result, val+1)
		} else {
			result = append(result, val)
		}
	}
	return result
}

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	idxArr := make([]int, 0, 2)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			idxArr = append(idxArr, i)
		}
		if len(idxArr) > 2 {
			return false
		}
	}
	switch len(idxArr) {
	case 0:
		return true
	case 1:
		return false
	case 2:
		return s1[idxArr[0]] == s2[idxArr[1]] && s1[idxArr[1]] == s2[idxArr[0]]
	}
	return false
}

func minAreaRect(points [][]int) int {
	m := make(map[int]map[int]bool)
	for _, v := range points {
		if _, ok := m[v[0]]; !ok {
			m[v[0]] = make(map[int]bool)
		}
		m[v[0]][v[1]] = true
	}
	min := math.MaxInt32
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i][0] == points[j][0] || points[i][1] == points[j][1] {
				continue
			}
			if _, ok := m[points[i][0]]; !ok {
				continue
			}
			if _, ok := m[points[j][0]]; !ok {
				continue
			}
			if _, ok := m[points[i][0]][points[j][1]]; !ok {
				continue
			}
			if _, ok := m[points[j][0]][points[i][1]]; !ok {
				continue
			}
			min = minInt(min, abs(points[i][0]-points[j][0])*abs(points[i][1]-points[j][1]))
		}
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func prisonAfterNDays(cells []int, n int) []int {
	state := uint8(0)
	for i, v := range cells {
		if v == 1 {
			state |= uint8(1) << uint(i)
		}
	}
	mask := (uint8(1) << (len(cells) - 1)) - 2
	state0 := state
	state = nextState(state0, mask)
	state1 := state
	loop := -1
	for i := 0; i < n-1; i++ {
		state = nextState(state, mask)
		if state == state1 {
			loop = i + 1
			break
		}
	}
	if loop > 0 {
		r := (n-1)%loop + 1
		state = state0
		for i := 0; i < r; i++ {
			state = nextState(state, mask)
		}
	}
	result := make([]int, len(cells))
	for i := 0; i < len(cells); i++ {
		result[i] = int(state>>uint(i)) & 1
	}
	return result
}

func nextState(state, mask uint8) uint8 {
	state = ^(state << 1) ^ (state >> 1)
	state &= mask
	return state
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	// return maxInt(maxConsecutiveChar(answerKey, 'T', k), maxConsecutiveChar(answerKey, 'F', k))
	return 0
}

// func maxConsecutiveChar(answerKey string, ch rune, k int) int {
// 	max := 0
// 	cnt := 0
// 	j := 0
// 	arr := []rune(answerKey)
// 	for i := range answerKey {
// 		if arr[i] != ch {
// 			cnt++
// 		}
// 		for cnt > k {
// 			if arr[j] != ch {
// 				cnt--
// 			}
// 			j++
// 		}
// 		max = maxInt(max, i-j+1)
// 	}
// 	return max
// }

func countTriples(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			tmpVal := i*i + j*j
			sqrt := math.Sqrt(float64(tmpVal))
			if sqrt == math.Trunc(sqrt) && int(sqrt) <= n {
				result += 2
			}
		}
	}
	return result
}

// func corpFlightBookings(bookings [][]int, n int) []int {
// 	result := make([]int, n)
// 	for _, v := range bookings {
// 		result[v[0]-1] += v[2]
// 		if v[1] < n {
// 			result[v[1]] -= v[2]
// 		}
// 	}
// 	for i := 1; i < n; i++ {
// 		result[i] += result[i-1]
// 	}
// 	return result
// }

func maxSumDivThree(nums []int) int {
	sum := 0
	inf := int(1e9)
	r1 := [2]int{inf, inf}
	r2 := [2]int{inf, inf}
	for _, v := range nums {
		sum += v
		r := v % 3
		if r == 1 {
			if v <= r1[0] {
				r1[1] = r1[0]
				r1[0] = v
			} else if v <= r1[1] {
				r1[1] = v
			}
		} else if r == 2 {
			if v <= r2[0] {
				r2[1] = r2[0]
				r2[0] = v
			} else if v <= r2[1] {
				r2[1] = v
			}
		}
	}
	result := 0
	r := sum % 3
	if r == 0 {
		result = sum
	}
	if r == 1 && r1[0] != inf {
		result = maxInt(result, sum-r1[0])
	}
	if r == 1 && r2[1] != inf {
		result = maxInt(result, sum-r2[0]-r2[1])
	}
	if r == 2 && r2[0] != inf {
		result = maxInt(result, sum-r2[0])
	}
	if r == 2 && r1[1] != inf {
		result = maxInt(result, sum-r1[0]-r1[1])
	}
	return result
}

// func maxInt(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
func busiestServers(k int, arrival, load []int) (ans []int) {
	available := hi{make([]int, k)}
	for i := 0; i < k; i++ {
		available.IntSlice[i] = i
	}
	busy := hp{}
	requests := make([]int, k)
	maxRequest := 0
	for i, t := range arrival {
		for len(busy) > 0 && busy[0].end <= t {
			heap.Push(&available, i+((busy[0].id-i)%k+k)%k) // 保证得到的是一个不小于 i 的且与 id 同余的数
			heap.Pop(&busy)
		}
		if available.Len() == 0 {
			continue
		}
		id := heap.Pop(&available).(int) % k
		requests[id]++
		if requests[id] > maxRequest {
			maxRequest = requests[id]
			ans = []int{id}
		} else if requests[id] == maxRequest {
			ans = append(ans, id)
		}
		heap.Push(&busy, pair{t + load[i], id})
	}
	return
}

type hi struct{ sort.IntSlice }

func (h *hi) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hi) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type pair struct{ end, id int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func canReorderDoubled(arr []int) bool {
	cnt := make(map[int]int, len(arr))
	for _, x := range arr {
		cnt[x]++
	}
	if cnt[0]%2 == 1 {
		return false
	}

	vals := make([]int, 0, len(cnt))
	for x := range cnt {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool { return abs(vals[i]) < abs(vals[j]) })

	for _, x := range vals {
		if cnt[2*x] < cnt[x] { // 无法找到足够的 2x 与 x 配对
			return false
		}
		cnt[2*x] -= cnt[x]
	}
	return true
}

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

func replaceDigits(s string) string {
	var sb strings.Builder
	for i, v := range s {
		if v >= '0' && v <= '9' {
			sb.WriteByte(s[i-1] + s[i] - '0')
		} else {
			sb.WriteByte(byte(v))
		}
	}
	return sb.String()
}

func totalFruit(fruits []int) int {
	basket := make(map[int]int)
	var (
		j      int
		result int
	)
	for i, v := range fruits {
		basket[v]++
		for len(basket) > 2 {
			basket[fruits[j]]--
			if basket[fruits[j]] == 0 {
				delete(basket, fruits[j])
			}
			j++
		}
		result = maxInt(result, i-j+1)
	}
	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func removeDuplicates(s string, k int) string {
	result := make([]byte, 0, len(s))
	cntArr := make([]int, len(s))
	for _, v := range s {
		result = append(result, byte(v))
		if len(result) <= 1 || byte(v) != result[len(result)-2] {
			cntArr[len(result)-1] = 1
			continue
		} else {
			cntArr[len(result)-1] = cntArr[len(result)-2] + 1
		}
		if cntArr[len(result)-1] == k {
			result = result[:len(result)-k]
			continue
		}
	}
	return string(result)
}

func strongPasswordChecker(password string) int {
	v0 := 0
	if len(password) > 20 {
		v0 = len(password) - 20
	}
	v1 := 6 - len(password)
	v2 := 0
	cnt := 1
	flags := [3]bool{false, false, false}
	for i := 1; i < len(password); i++ {
		if 'a' <= password[i] && password[i] <= 'z' {
			flags[0] = true
		}
		if 'A' <= password[i] && password[i] <= 'Z' {
			flags[1] = true
		}
		if '0' <= password[i] && password[i] <= '9' {
			flags[2] = true
		}
		if password[i] == password[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt >= 3 {
			v2 += cnt - 2
			cnt = 0
		}
	}
	v3 := 0
	for _, v := range flags {
		if !v {
			v3++
		}
	}

	return v0 + maxInt(v1, maxInt(v2, v3))
}

func divideArray(nums []int) bool {
	cntMap := make(map[int]int, 0)
	for _, v := range nums {
		cntMap[v]++
	}
	for _, v := range cntMap {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func minBitFlips(start int, goal int) int {
	diff := goal ^ start
	cnt := 0
	for diff != 0 {
		if diff&1 == 1 {
			cnt++
		}
		diff >>= 1
	}
	return cnt
}

func triangularSum(nums []int) int {
	cur := nums
	for i := 0; i < len(nums)-1; i++ {
		temp := make([]int, len(cur)-1)
		for j := 0; j < len(temp); j++ {
			temp[j] = (cur[j] + cur[j+1]) % 10
		}
		cur = temp
	}
	return cur[0]
}

func numberOfWays(s string) int64 {
	cnt := int64(0)
	result := int64(0)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cnt++
		}
	}
	cur := int64(0)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			result += (int64(i) - cur) * (int64(len(s)) - cnt - int64(i) + cur)
			cur++
		} else {
			result += cur * int64(cnt-cur)
		}
	}
	return result
}

func sumScores(s string) int64 {
	result := int64(len(s))
	arr := []byte(s)
	for i := 1; i < len(arr); i++ {
		sub := arr[i:]
		for j := 0; j < len(sub); j++ {
			if arr[j] != sub[j] {
				break
			}
			result++
		}
	}
	return result
}

func firstPalindrome(words []string) string {
	for _, v := range words {
		if isPalindrome(v) {
			return v
		}
	}
	return ""
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// Generated by https://quicktype.io

type Person struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
}

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// func nextLargerNodes(head *ListNode) []int {
// 	var hn NodeArr
// 	length := Len(head)
// 	result := make([]int, length)
// 	for i := 0; i < length; i++ {
// 		for hn.Len() > 0 && hn[0].Val < head.Val {
// 			node := heap.Pop(&hn).(ListNodeIndex)
// 			result[node.Index] = head.Val
// 		}
// 		heap.Push(&hn, ListNodeIndex{Index: i, ListNode: *head})
// 		head = head.Next
// 	}
// 	return result
// }

// func Len(head *ListNode) int {
// 	cnt := 0
// 	for head != nil {
// 		cnt++
// 		head = head.Next
// 	}
// 	return cnt
// }

// type ListNodeIndex struct {
// 	ListNode
// 	Index int
// }

// type NodeArr []ListNodeIndex

// func (n NodeArr) Len() int {
// 	return len(n)
// }

// func (n NodeArr) Less(i, j int) bool {
// 	return n[i].Val < n[j].Val
// }

// func (n NodeArr) Swap(i, j int) {
// 	n[i], n[j] = n[j], n[i]
// }

// func (n *NodeArr) Push(x interface{}) {
// 	*n = append(*n, x.(ListNodeIndex))
// }

// func (n *NodeArr) Pop() interface{} {
// 	old := *n
// 	x := old[len(old)-1]
// 	*n = old[:len(old)-1]
// 	return x
// }

func nextLargerNodes(head *ListNode) []int {
	stack := make([]ListNodeIndex, 0)
	length := Len(head)
	result := make([]int, length)
	for i := 0; i < length; i++ {
		for len(stack) > 0 && stack[len(stack)-1].Val < head.Val {
			result[stack[len(stack)-1].Index] = head.Val
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, ListNodeIndex{Index: i, ListNode: *head})
		head = head.Next
	}

	return result
}

type ListNodeIndex struct {
	ListNode
	Index int
}

func Len(head *ListNode) int {
	cnt := 0
	for head != nil {
		cnt++
		head = head.Next
	}
	return cnt
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	found := true
	for found {
		cur := dummy.Next
		sumMap := make(map[int]*ListNode, 0)
		sumMap[0] = dummy
		sum := 0
		found = false
		for cur != nil {
			sum += cur.Val
			if node, ok := sumMap[sum]; ok {
				node.Next = cur.Next
				found = true
				break
			}
			sumMap[sum] = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}

// It returns the maximum absolute value of the difference between any two elements in the array.
func maxAbsValExpr(arr1 []int, arr2 []int) int {
	inf := int(1e7)
	maxArr := [4]int{-inf, -inf, -inf, -inf}
	minArr := [4]int{inf, inf, inf, inf}

	for i := 0; i < len(arr1); i++ {
		maxArr[0] = max(maxArr[0], arr1[i]+arr2[i]+i)
		minArr[0] = min(minArr[0], arr1[i]+arr2[i]+i)
		maxArr[1] = max(maxArr[1], arr1[i]-arr2[i]+i)
		minArr[1] = min(minArr[1], arr1[i]-arr2[i]+i)
		maxArr[2] = max(maxArr[2], -arr1[i]+arr2[i]+i)
		minArr[2] = min(minArr[2], -arr1[i]+arr2[i]+i)
		maxArr[3] = max(maxArr[3], -arr1[i]-arr2[i]+i)
		minArr[3] = min(minArr[3], -arr1[i]-arr2[i]+i)
	}
	result := 0
	for i := 0; i < 4; i++ {
		result = max(result, maxArr[i]-minArr[i])
	}
	return result
}

// Unzip a file
func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// zip a dir
// srcFile could be a single file or a directory
func Zip(srcFile string, destZip string) error {
	zipfile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = path
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			io.Copy(writer, file)
		}
		return err
	})

	return err
}

// > Tar takes a source file and a destination tar file and returns an error
func Tar(srcFile string, destTar string) error {
	file, err := os.Create(destTar)
	if err != nil {
		return err
	}
	defer file.Close()

	tw := tar.NewWriter(file)
	defer tw.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		if info.IsDir() {
			header.Name += "/"
		}

		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			io.Copy(tw, file)
		}
		return err
	})

	return err
}

func UnTar(srcTar string, destDir string) error {
	file, err := os.Open(srcTar)
	if err != nil {
		return err
	}
	defer file.Close()

	tr := tar.NewReader(file)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(destDir, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			os.MkdirAll(path, info.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), os.ModePerm)
			file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(file, tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// create http server
func CreateHttpServer(addr string, handler http.Handler) error {
	server := &http.Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

// parse xml file
func ParseXml(file string, v interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, v)
}

// bean to xml file
func BeanToXml(bean interface{}, file string) error {
	data, err := xml.MarshalIndent(bean, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, data, 0644)
}

// generate random string
func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

// draw line chart with go-chart
// func DrawLineChart(title string, xAxis []string, yAxis []float64, file string) error {
// 	chart := chart.Chart{
// 		Title: title,
// 		XAxis: chart.XAxis{
// 			Style: chart.StyleShow(),
// 			Data:  xAxis,
// 		},
// 		YAxis: chart.YAxis{
// 			Style: chart.StyleShow(),
// 		},
// 		Series: []chart.Series{
// 			chart.ContinuousSeries{
// 				Style: chart.Style{
// 					Show:        true,
// 					StrokeColor: chart.ColorBlue,
// 					FillColor:   chart.ColorBlue.WithAlpha(64),
// 				},
// 				XValues: xAxis,
// 				YValues: yAxis,
// 			},
// 		},
// 	}
// 	return chart.Save(file, chart.PNG)
// }

func sumBase(n int, k int) int {
	result := 0
	for n != 0 {
		result += n % k
		n /= k
	}
	fmt.Println("result:", result)
	return result
}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)

	var preorder func(node *Node, level int)
	preorder = func(node *Node, level int) {
		if node == nil {
			return
		}
		if len(result) == level {
			result = append(result, []int{})
		}
		result[level] = append(result[level], node.Val)
		for _, child := range node.Children {
			preorder(child, level+1)
		}
	}
	preorder(root, 0)
	return result
}

func countQuadruplets(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						result++
					}
				}
			}
		}
	}
	return result
}

func numSpecial(mat [][]int) int {
	var result int
	rowCnt := make([]int, len(mat))
	colCnt := make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[1]); j++ {
			if mat[i][j] == 1 {
				rowCnt[i]++
				colCnt[j]++
			}
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 && rowCnt[i] == 1 && colCnt[j] == 1 {
				result++
			}
		}
	}
	return result
}

func minAreaFreeRect(points [][]int) float64 {
	var result float64 = math.MaxFloat64
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				for l := k + 1; l < len(points); l++ {
					if isRectangle(points[i], points[j], points[k], points[l]) {
						result = minFloat(result, area(points[i], points[j], points[k], points[l]))
					}
				}
			}
		}
	}
	if result == math.MaxFloat64 {
		return 0
	}
	return result
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func isRectangle(p1, p2, p3, p4 []int) bool {
	ps := [][]int{p1, p2, p3, p4}
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			for k := j + 1; k < len(ps); k++ {
				v1 := (ps[i][0]-ps[j][0])*(ps[i][0]-ps[j][0]) + (ps[i][1]-ps[j][1])*(ps[i][1]-ps[j][1])
				v2 := (ps[i][0]-ps[k][0])*(ps[i][0]-ps[k][0]) + (ps[i][1]-ps[k][1])*(ps[i][1]-ps[k][1])
				v3 := (ps[j][0]-ps[k][0])*(ps[j][0]-ps[k][0]) + (ps[j][1]-ps[k][1])*(ps[j][1]-ps[k][1])
				if v1+v2 == v3 || v1+v3 == v2 || v2+v3 == v1 {
					continue
				}
				return false
			}
		}
	}
	return true
}

func area(p1, p2, p3, p4 []int) float64 {
	v1 := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
	v2 := (p1[0]-p3[0])*(p1[0]-p3[0]) + (p1[1]-p3[1])*(p1[1]-p3[1])
	v3 := (p3[0]-p2[0])*(p3[0]-p2[0]) + (p3[1]-p2[1])*(p3[1]-p2[1])
	vs := []int{v1, v2, v3}
	sort.Ints(vs)
	return math.Sqrt(float64(vs[0])) * math.Sqrt(float64(vs[1]))
}

func numberOfLines(widths []int, s string) []int {
	var result []int = make([]int, 2)
	result[0] = 1
	cnt := 0
	for _, v := range s {
		cnt += widths[v-'a']
		if cnt > 100 {
			result[0]++
			cnt = widths[v-'a']
		}
	}
	result[1] = cnt
	return result
}

// func groupThePeople(groupSizes []int) [][]int {
// 	var result [][]int = make([][]int, 0)
// 	sizeMap := make(map[int][]int, 0)
// 	for i := 0; i < len(groupSizes); i++ {
// 		group, ok := sizeMap[groupSizes[i]]
// 		if !ok {
// 			group = make([]int, 0, groupSizes[i])
// 			sizeMap[groupSizes[i]] = group
// 		}
// 		group = append(group, i)
// 		sizeMap[groupSizes[i]] = group
// 		if len(group) == groupSizes[i] {
// 			result = append(result, group)
// 			delete(sizeMap, groupSizes[i])
// 		}
// 	}
// 	return result
// }

func makeEqual(words []string) bool {
	cntArr := make([]int, 26)
	for _, v := range words {
		for _, vv := range v {
			cntArr[vv-'a']++
		}
	}
	for _, v := range cntArr {
		if v%len(words) != 0 {
			return false
		}
	}
	return true
}

type RandomizedSet struct {
	data map[int]int
	keys []int
	rand *rand.Rand
}

// func Constructor() RandomizedSet {
// 	return RandomizedSet{
// 		data: make(map[int]int, 0),
// 		keys: make([]int, 0),
// 		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
// 	}
// }

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.data[val]; ok {
		return false
	}
	this.data[val] = len(this.keys)
	this.keys = append(this.keys, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.data[val]; !ok {
		return false
	}
	idx := this.data[val]
	last := len(this.keys) - 1
	this.keys[idx] = this.keys[last]
	this.data[this.keys[idx]] = idx
	this.keys = this.keys[:last]
	delete(this.data, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.keys[this.rand.Intn(len(this.keys))]
}

func networkDelayTime(times [][]int, n int, k int) int {
	tree := make(map[int]map[int]int, 0)
	for _, v := range times {
		if _, ok := tree[v[0]-1]; !ok {
			tree[v[0]-1] = make(map[int]int, 0)
		}
		tree[v[0]-1][v[1]-1] = v[2]
	}
	dist := make([]int, n)
	visited := make([]bool, n)
	inf := int(1e6)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}
	dist[k-1] = 0
	for i := 0; i < n; i++ {
		min := -1
		for i := 0; i < n; i++ {
			if !visited[i] && (min == -1 || dist[i] < dist[min]) {
				min = i
			}
		}
		visited[min] = true
		for j, v := range tree[min] {
			if dist[min]+v < dist[j] {
				dist[j] = dist[min] + v
			}
		}
	}

	max := -1
	for i := 0; i < n; i++ {
		if max == -1 || dist[i] > dist[max] {
			max = i
		}
	}
	if dist[max] == inf {
		return -1
	}
	return dist[max]
}

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	upper := int(math.Pow10(n)) - 1
	for left := upper; ; left-- {
		p := left
		for x := left; x > 0; x /= 10 {
			p = p*10 + x%10
		}
		for x := upper; x*x >= p; x-- {
			if p%x == 0 { // x 是 p 的因子
				return p % 1337
			}
		}
	}
	return -1
}

func isPalindromeInt(val int64) bool {
	arr := make([]int64, 0)
	for val != 0 {
		arr = append(arr, val%10)
		val /= 10
	}
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}

func medianSlidingWindow(nums []int, k int) []float64 {
	var result []float64 = make([]float64, 0)
	if len(nums) == 0 {
		return result
	}
	return result
}

// type mhp struct {
// 	sort.IntSlice
// }

// func (he *mhp) Push(x interface{}) {
// 	he.IntSlice = append(he.IntSlice, x.(int))
// }

// func (he *mhp) Pop() interface{} {
// 	n := len(he.IntSlice)
// 	x := he.IntSlice[n-1]
// 	he.IntSlice = he.IntSlice[:n-1]
// 	return x
// }

func largest1BorderedSquare(grid [][]int) int {
	rowCnt := make([][]int, len(grid))
	colCnt := make([][]int, len(grid))
	result := 0
	for i := 0; i < len(grid); i++ {
		rowCnt[i] = make([]int, len(grid[0]))
		colCnt[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 {
				rowCnt[i][j] = grid[i][j]
			} else if grid[i][j] == 1 {
				rowCnt[i][j] = rowCnt[i-1][j] + grid[i][j]
			}
			if j == 0 {
				colCnt[i][j] = grid[i][j]
			} else if grid[i][j] == 1 {
				colCnt[i][j] = colCnt[i][j-1] + grid[i][j]
			}

			min(rowCnt[i][j]-1, colCnt[i][j]-1)
			for k := min(rowCnt[i][j]-1, colCnt[i][j]-1); k >= 0; k-- {
				if rowCnt[i][j-k] >= k+1 && colCnt[i-k][j] >= k+1 {
					if k+1 > result {
						result = k + 1
					}
					break
				}
			}
		}
	}
	// fmt.Println(rowCnt, colCnt)
	return result * result
}

// convert bytes to string
func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func bytesToString2(b []byte) string {
	return string(b)
}
