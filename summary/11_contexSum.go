package main

import (
	"context"
	"fmt"
	"time"
)

// 单独的监控协程
func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "收到信号，监控退出,time=", time.Now().Unix())
			return
		default:
			fmt.Println(name, "goroutine监控中,time=", time.Now().Unix())
			time.Sleep(1 * time.Second)
		}
	}
}

func demoWithDeadline() {
	d := time.Now().Add(2*time.Second - 100*time.Millisecond)

	// 接受一个 Context 和过期时间作为参数，返回其子Context和取消函数cancel
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	go watch(ctx, "Deadline watch_1")
	go watch(ctx, "Deadline watch_2")

	time.Sleep(5 * time.Second)
}

func demoWithCancel() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func demoWithTimeout() {
	timeout := 2*time.Second - 100*time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go watch(ctx, "Timeout watch_1")
	go watch(ctx, "Timeout watch_2")

	time.Sleep(5 * time.Second)
}

func demoWithValue() {
	type contextKey string

	ctx := context.Background()
	ctxA := context.WithValue(ctx, contextKey("keyA"), "valueA")
	ctxB := context.WithValue(ctxA, contextKey("keyB"), "valueB")

	f := func(ctx context.Context, k contextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found key value:", k, v)
			return
		}
		fmt.Println("key not found:", k)
	}

	f(ctxA, contextKey("keyA"))
	f(ctxA, contextKey("keyB"))
	f(ctxB, contextKey("keyA"))
	f(ctxB, contextKey("keyB"))
}

// func main() {
// 	demoWithCancel()
// 	demoWithDeadline()
// 	demoWithTimeout()
// 	demoWithValue()
// }
