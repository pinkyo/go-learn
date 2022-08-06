package main

import "fmt"

func main() {
	// fmt.Println(calculateTax([][]int{
	// 	{3, 50}, {7, 10}, {12, 25},
	// }, 10))
	// fmt.Println("Hello World!")
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i+1, "hello world")
	// }
	// fmt.Println("Hello World!")
	// fmt.Println("Hello World!")
	// fmt.Println("ls -al")
	// fmt.Println("Hello World!")
	fmt.Println("Hello World!")
}

func calculateTax(brackets [][]int, income int) float64 {
	var result float64
	pre := 0
	for i := 0; i < len(brackets); i++ {
		bracket := brackets[i]
		if bracket[0] >= income {
			result += float64(income-pre) * float64(bracket[1])
			return result / 100
		}
		result += float64(bracket[0]-pre) * float64(bracket[1])
		pre = bracket[0]
	}
	return result / 100
}
