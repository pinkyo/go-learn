package main

import "fmt"

func main() {
	fmt.Println(string(findKthBit(3, 1)))
	fmt.Println(string(findKthBit(4, 12)))
}

// func findKthBit(n int, k int) byte {
// 	str := "0"
// 	for i := 1; i < n; i++ {
// 		str = str + "1" + reverse(invert(str))
// 		// fmt.Println("str: ", str)
// 	}
// 	return str[k-1]
// }

// func invert(str string) string {
// 	var result string
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == '1' {
// 			result += "0"
// 		} else {
// 			result += "1"
// 		}
// 	}
// 	return result
// }

// func reverse(str string) string {
// 	var result string
// 	for i := len(str) - 1; i >= 0; i-- {
// 		result += string(str[i])
// 	}
// 	return result
// }

// S4 = 011100110110001
func findKthBit(n int, k int) byte {
	var lenArr []int
	lenArr = append(lenArr, 1)
	for i := 1; i < n; i++ {
		lenArr = append(lenArr, 2*lenArr[i-1]+1)
	}
	var cnt int
	for i := len(lenArr) - 1; i > 0; i-- {
		if k > lenArr[i-1]+1 {
			cnt++
			k = lenArr[i] - k + 1
		} else if k == lenArr[i-1]+1 {
			return '1' - byte(cnt%2)
		}
	}
	return '0' + byte(cnt%2)
}
