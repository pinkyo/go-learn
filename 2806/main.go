package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	fmt.Println(accountBalanceAfterPurchase(15))
	fmt.Println(accountBalanceAfterPurchase(14))
}

func accountBalanceAfterPurchase(purchaseAmount int) int {
	remain := purchaseAmount % 10
	if remain >= 5 {
		purchaseAmount -= remain
		purchaseAmount += 10
	} else {
		purchaseAmount -= remain
	}
	return 100 - purchaseAmount
}
