package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync"
	"time"
)

// ============================================================
// 版本1：最基础的 SafeGo - 无参数函数，自动 recover panic
// ============================================================

func SafeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[SafeGo] panic recovered: %v\n%s", r, debug.Stack())
			}
		}()
		fn()
	}()
}

// ============================================================
// 版本2：带 PanicHandler 的 SafeGo - 可自定义 panic 处理逻辑
// ============================================================

type PanicHandler func(r any, stack []byte)

func SafeGoWithHandler(fn func(), handler PanicHandler) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if handler != nil {
					handler(r, debug.Stack())
				} else {
					log.Printf("[SafeGo] panic recovered: %v\n%s", r, debug.Stack())
				}
			}
		}()
		fn()
	}()
}

// ============================================================
// 版本3：使用泛型支持单个参数 - Go 1.18+
// ============================================================

func SafeGo1[T any](fn func(T), arg T) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[SafeGo1] panic recovered: %v\n%s", r, debug.Stack())
			}
		}()
		fn(arg)
	}()
}

// 两个参数版本
func SafeGo2[T1, T2 any](fn func(T1, T2), arg1 T1, arg2 T2) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[SafeGo2] panic recovered: %v\n%s", r, debug.Stack())
			}
		}()
		fn(arg1, arg2)
	}()
}

// 三个参数版本
func SafeGo3[T1, T2, T3 any](fn func(T1, T2, T3), arg1 T1, arg2 T2, arg3 T3) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[SafeGo3] panic recovered: %v\n%s", r, debug.Stack())
			}
		}()
		fn(arg1, arg2, arg3)
	}()
}

// ============================================================
// 版本4：支持可变参数（不推荐，类型不安全，但灵活）
// ============================================================

func SafeGoAny(fn func(args ...any), args ...any) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[SafeGoAny] panic recovered: %v\n%s", r, debug.Stack())
			}
		}()
		fn(args...)
	}()
}

// ============================================================
// 版本5：返回 error 通道 - 可以等待 goroutine 完成并获取 panic 错误
// ============================================================

func SafeGoErr(fn func()) chan error {
	errCh := make(chan error, 1)
	go func() {
		defer close(errCh)
		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()
				errCh <- fmt.Errorf("panic: %v\nstack: %s", r, stack)
			}
		}()
		fn()
	}()
	return errCh
}

// 带参数且返回 error
func SafeGoErr1[T any](fn func(T), arg T) chan error {
	errCh := make(chan error, 1)
	go func() {
		defer close(errCh)
		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()
				errCh <- fmt.Errorf("panic: %v\nstack: %s", r, stack)
			}
		}()
		fn(arg)
	}()
	return errCh
}

// ============================================================
// 版本6：更完善的方案 - 使用闭包捕获参数（最推荐，灵活且类型安全）
// ============================================================

// Go 函数是一等公民，闭包天然支持捕获外部变量，这是最地道的 Go 风格实现

// SafeGoFunc 是一个通用的安全 goroutine 启动器
// 使用方式: SafeGoFunc(func() { yourFunc(a, b, c) })
func SafeGoFunc(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[SafeGoFunc] panic recovered: %v\n%s", r, debug.Stack())
			}
		}()
		fn()
	}()
}

// ============================================================
// 版本7：带 WaitGroup 支持，可等待完成
// ============================================================

func SafeGoWg(wg *sync.WaitGroup, fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[SafeGoWg] panic recovered: %v\n%s", r, debug.Stack())
			}
		}()
		fn()
	}()
}

// ============================================================
// 版本8：生产级完整实现 - 可配置的 SafeGo
// ============================================================

type SafeGoOption func(*safeGoConfig)

type safeGoConfig struct {
	handler      PanicHandler
	waitGroup  *sync.WaitGroup
	errChan    chan error
}

func WithPanicHandler(h PanicHandler) SafeGoOption {
	return func(c *safeGoConfig) {
		c.handler = h
	}
}

func WithWaitGroup(wg *sync.WaitGroup) SafeGoOption {
	return func(c *safeGoConfig) {
		c.waitGroup = wg
	}
}

func WithErrChan(ch chan error) SafeGoOption {
	return func(c *safeGoConfig) {
		c.errChan = ch
	}
}

func SafeGoPro(fn func(), opts ...SafeGoOption) {
	cfg := &safeGoConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.waitGroup != nil {
		cfg.waitGroup.Add(1)
	}

	go func() {
		if cfg.waitGroup != nil {
			defer cfg.waitGroup.Done()
		}

		defer func() {
			if r := recover(); r != nil {
				stack := debug.Stack()
				if cfg.handler != nil {
					cfg.handler(r, stack)
				} else {
					log.Printf("[SafeGoPro] panic recovered: %v\n%s", r, stack)
				}
				if cfg.errChan != nil {
					cfg.errChan <- fmt.Errorf("panic: %v", r)
				}
			}
		}()

		fn()
	}()
}

// ============================================================
// 示例：如何使用
// ============================================================

func demoPanic(msg string) {
	panic(msg)
}

func greet(name string, age int) {
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

func riskyPrint(items ...any) {
	for _, item := range items {
		fmt.Printf("item: %v\n", item)
	}
	panic("oops in riskyPrint")
}

func main() {
	fmt.Println("===== 版本1：基础 SafeGo =====")
	SafeGo(func() {
		fmt.Println("basic safe goroutine")
	})
	SafeGo(func() {
		demoPanic("panic in basic")
	})
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n===== 版本2：带自定义 PanicHandler =====")
	SafeGoWithHandler(
		func() { demoPanic("custom handler panic") },
		func(r any, stack []byte) {
			fmt.Printf("CUSTOM HANDLER: got panic=%v\n", r)
		},
	)
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n===== 版本3：泛型参数 SafeGo1/SafeGo2 =====")
	SafeGo1(func(name string) {
		fmt.Printf("Hello, %s!\n", name)
		panic("panic in SafeGo1")
	}, "Alice")

	SafeGo2(greet, "Bob", 30)
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n===== 版本4：可变参数 SafeGoAny（不推荐） =====")
	SafeGoAny(riskyPrint, "a", 1, true)
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n===== 版本5：返回 error 通道 =====")
	errCh := SafeGoErr(func() {
		demoPanic("panic in Err version")
	})
	if err := <-errCh; err != nil {
		fmt.Println("Got error from errCh:", err)
	}

	fmt.Println("\n===== 版本6：闭包方式（最推荐） =====")
	// 使用闭包传参，类型安全且任意参数
	name := "Charlie"
	score := 95
	SafeGoFunc(func() {
		fmt.Printf("%s scored %d points\n", name, score)
		panic("panic in closure")
	})
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n===== 版本7：WaitGroup 支持 =====")
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		idx := i
		SafeGoWg(&wg, func() {
			fmt.Printf("wg task %d running\n", idx)
			if idx == 1 {
				panic(fmt.Sprintf("panic in wg task %d", idx))
			}
		})
	}
	wg.Wait()
	fmt.Println("all wg tasks done (including panicked ones also marked done)")

	fmt.Println("\n===== 版本8：生产级 SafeGoPro =====")
	var wg2 sync.WaitGroup
	errCh2 := make(chan error, 2)

	for i := 0; i < 2; i++ {
		idx := i
		SafeGoPro(
			func() {
				fmt.Printf("pro task %d running\n", idx)
				if idx == 0 {
					panic(fmt.Sprintf("panic in pro task %d", idx))
				}
			},
			WithWaitGroup(&wg2),
			WithErrChan(errCh2),
			WithPanicHandler(func(r any, stack []byte) {
				fmt.Printf("PRO HANDLER: panic=%v\n", r)
			}),
		)
	}
	wg2.Wait()
	close(errCh2)
	for err := range errCh2 {
		fmt.Println("pro err:", err)
	}

	fmt.Println("\n===== 所有示例执行完毕 =====")
}