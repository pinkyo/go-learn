package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 方式一：使用 context.WithCancel 通知 goroutine 退出
func withContext() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done(): // 收到取消信号
				fmt.Println("[ctx] goroutine 收到关闭信号，执行清理后退出")
				return
			default:
				fmt.Println("[ctx] working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel()  // 发送关闭信号
	wg.Wait() // 等待 goroutine 真正退出
}

// 方式二：使用专用的 done channel 通知 goroutine 退出
func withChannel() {
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done: // done 被关闭后，所有监听者都能收到信号
				fmt.Println("[chan] goroutine 收到关闭信号，执行清理后退出")
				return
			default:
				fmt.Println("[chan] working...")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	close(done) // 关闭通道，广播关闭信号
	wg.Wait()   // 等待 goroutine 真正退出
}

func main() {
	withContext()
	fmt.Println("---")
	withChannel()
}
