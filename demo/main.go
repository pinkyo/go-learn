package main

import "fmt"

func main() {
	// ch1 := make(chan int, 1)
	// ch2 := make(chan int, 1)

	// var wg sync.WaitGroup
	// wg.Add(2)

	// ch1 <- 1
	// go func() {
	// 	defer wg.Done()
	// 	for i := range ch1 {
	// 		fmt.Println(i)
	// 		ch2 <- i + 1
	// 	}
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	for i := range ch2 {
	// 		fmt.Println(i)
	// 		if i == 100 {
	// 			close(ch1)
	// 			close(ch2)
	// 		} else {
	// 			ch1 <- i + 1
	// 		}
	// 	}
	// }()

	// wg.Wait()

	arr := [10]int{}
	fmt.Println(arr)

	slice := arr[0:5]
	slice[0] = 1000
	fmt.Println(arr)
	fmt.Println(slice)
}
