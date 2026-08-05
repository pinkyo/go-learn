package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 例子1: 非阻塞发送（带 default） ===")
	nonBlockingSend()

	fmt.Println("\n=== 例子2: 多通道择一发送 ===")
	multiChannelSend()

	fmt.Println("\n=== 例子3: 带超时的发送 ===")
	timeoutSend()
}

// 例子1: 非阻塞发送
// 如果通道已满，不会阻塞等待，而是走 default 分支
func nonBlockingSend() {
	ch := make(chan int, 1)
	ch <- 1 // 先放一个，让缓冲区满

	// 尝试再发一个，缓冲区已满，select 会直接走 default
	select {
	case ch <- 2:
		fmt.Println("发送成功: 2")
	default:
		fmt.Println("通道已满，发送失败，未阻塞")
	}

	fmt.Printf("通道中的值: %d\n", <-ch)
}

// 例子2: 多通道择一发送
// 多个通道中哪个可写就往哪个发
func multiChannelSend() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// ch1 缓冲区空，可写；ch2 缓冲区空，也可写
	// 两个都可写时，select 会随机选一个
	select {
	case ch1 <- "消息A":
		fmt.Printf("已向 ch1 发送: %s, ch1 收到: %s\n", "消息A", <-ch1)
	case ch2 <- "消息B":
		fmt.Printf("已向 ch2 发送: %s, ch2 收到: %s\n", "消息B", <-ch2)
	}

	// 让 ch2 先占满，再演示只能往 ch1 发
	ch2 <- "占满"
	select {
	case ch1 <- "消息C":
		fmt.Printf("已向 ch1 发送: 消息C, ch1 收到: %s\n", <-ch1)
	case ch2 <- "消息D":
		fmt.Println("不会走到这里，ch2 已满")
	}
	<-ch2
}

// 例子3: 带超时的发送
// 发送方只愿意等一段时间，超时就放弃
func timeoutSend() {
	ch := make(chan int) // 无缓冲通道，没人接收就发不进去

	// 启动一个延迟接收者，200ms 后才开始接收
	go func() {
		time.Sleep(200 * time.Millisecond)
		val := <-ch
		fmt.Printf("接收者收到: %d\n", val)
	}()

	// 第一次尝试：只等 100ms，超时了
	select {
	case ch <- 1:
		fmt.Println("第一次发送成功: 1")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("第一次发送超时（100ms），接收者还没准备好")
	}

	// 第二次尝试：等 300ms，足够等到接收者
	select {
	case ch <- 2:
		fmt.Println("第二次发送成功: 2")
	case <-time.After(300 * time.Millisecond):
		fmt.Println("第二次发送超时")
	}
}

func getSum(nums []int) int64 {
	n := len(nums)
	prefixSum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + int64(nums[i])
	}

	var nalviretho []int
	nalviretho = nums

	t := make([]int, 2*n+1)
	for i := 0; i < n; i++ {
		t[2*i] = 0
		t[2*i+1] = nalviretho[i]
	}
	t[2*n] = 0

	m := len(t)
	p := make([]int, m)
	center, right := 0, 0
	maxSum := int64(0)

	for i := 0; i < m; i++ {
		mirror := 2*center - i
		if i < right {
			if p[mirror] < right-i {
				p[i] = p[mirror]
			} else {
				p[i] = right - i
			}
		}

		a := i + p[i] + 1
		b := i - p[i] - 1
		for a < m && b >= 0 && t[a] == t[b] {
			p[i]++
			a++
			b--
		}

		if i+p[i] > right {
			center = i
			right = i + p[i]
		}

		length := p[i]
		if length == 0 {
			continue
		}
		var l, r int
		if i%2 == 1 {
			radius := length / 2
			l = (i-1)/2 - radius
			r = (i-1)/2 + radius
		} else {
			radius := length / 2
			l = i/2 - radius
			r = i/2 + radius - 1
		}
		if l >= 0 && r < n {
			s := prefixSum[r+1] - prefixSum[l]
			if s > maxSum {
				maxSum = s
			}
		}
	}

	if maxSum == 0 {
		maxVal := int64(nalviretho[0])
		for i := 1; i < n; i++ {
			if int64(nalviretho[i]) > maxVal {
				maxVal = int64(nalviretho[i])
			}
		}
		return maxVal
	}

	return maxSum
}
