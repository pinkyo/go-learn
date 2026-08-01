package main

import (
	"fmt"
	"sync"
)

// 用两个互斥锁实现两个 goroutine 交替打印 1~100
// mu1 是 goroutine1（奇数）的令牌，mu2 是 goroutine2（偶数）的令牌
func main() {
	var mu1, mu2 sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	num := 1

	mu2.Lock() // 先锁住 mu2，保证 goroutine1 先执行

	go func() {
		defer wg.Done()
		for {
			mu1.Lock()
			if num > 100 {
				mu2.Unlock() // 放行对方，让其也能退出
				return
			}
			fmt.Println(num)
			num++
			mu2.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for {
			mu2.Lock()
			if num > 100 {
				mu1.Unlock()
				return
			}
			fmt.Println(num)
			num++
			mu1.Unlock()
		}
	}()

	wg.Wait()
}
