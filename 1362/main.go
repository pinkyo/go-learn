package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(closestDivisors(999))
	fmt.Println(closestDivisors(8))
	fmt.Println(closestDivisors(123))
}

func closestDivisors(num int) []int {
	from := int(math.Floor(math.Sqrt(float64(num + 2))))
	for i := from; i >= 0; i-- {
		if (num+1)%i == 0 {
			return []int{i, (num + 1) / i}
		}
		if (num+2)%i == 0 {
			return []int{i, (num + 2) / i}
		}
	}
	return []int{} // never reach here
}
