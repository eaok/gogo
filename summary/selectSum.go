package main

import (
	"fmt"
	"sync"
	"time"
)

func baseSelect() {
	ch := make(chan int, 1)

	ch <- 1

	//多个chan有数据时，会随机选择执行，都没数据时才会执行default
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	default:
		fmt.Println("exit")
	}
}

//Timeout 超时机制
func timeoutSelect() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	ch := make(chan int)

	//1.通过timeout通道来控制超时
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("timeout 01")
	}

	//2.通过time.After来控制超时
	select {
	case <-ch:
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 02")
	}
}

//检查 channel 是否已满
func checkChanFullSelect() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	select {
	case ch <- 2:
		fmt.Println("channel value is", <-ch)
		fmt.Println("channel value is", <-ch)
	default:
		fmt.Println("channel blocking")
	}
}

//select for loop 用法
func loopSelect() {
	var wg sync.WaitGroup
	ch := make(chan int, 4)
	out := make(chan string, 0)

	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 9; i++ {
			ch <- i
		}
		time.Sleep(2 * time.Second)
		out <- "stop"
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
	LOOP:
		for {
			select {
			case m := <-ch:
				println(m)
			case <-out:
				break LOOP
			default:
			}
		}
		wg.Done()
	}(&wg)

	wg.Wait()
}

func main() {
	baseSelect()
	timeoutSelect()
	checkChanFullSelect()
	loopSelect()
}
