package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// chanSend
// 如果Channel的 recvq 上有等待的 Goroutine，就会向chanel发送数据;
func chanWaitSend() {
	ch := make(chan int)

	go func(ch chan int) {
		for v := range ch {
			time.Sleep(time.Second)
			fmt.Println("recv ", v)
		}
	}(ch)

	for i := 0; i < 4; i++ {
		select {
		case ch <- i:
			fmt.Println("send ", i)
		}
	}

	time.Sleep(time.Second * 2)
}

// chanBufferSend
// 如果Channel的缓冲区存在空闲位置，就会将待发送的数据存入缓冲区，否则阻塞等待；
func chanBufferSend() {
	chBuffer := make(chan int, 3)
	go func(ch chan int) {
		for v := range ch {
			time.Sleep(time.Second)
			fmt.Println("recv ", v)
		}
	}(chBuffer)

	for i := 0; i < 9; i++ {
		select {
		case chBuffer <- i:
			fmt.Printf("buffer=%d send %d\n", len(chBuffer), i)
		}
	}
}

func caseSend() {
	//chanWaitSend()
	chanBufferSend()
}

// chanWaitRecv
// 如果当前 Channel 的 sendq 上有等待的 Goroutine，就会读取数据
func chanWaitRecv() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 2)
			ch <- i
			fmt.Println("send ", i)
		}
		close(ch)
		done <- true
	}()

LOOP:
	for {
		select {
		case data := <-ch:
			fmt.Println("recv ", data)
		case <-done:
			break LOOP
		}
	}
}

// chanBufferRecv
// 如果当前 Channel 的缓冲区不为空，也会读取数据
func chanBufferRecv() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 9; i++ {
			ch <- i
			fmt.Println("send ", i)
		}
		close(ch)
	}()

LOOP:
	for {
		time.Sleep(time.Second)
		select {
		case data := <-ch:
			fmt.Printf("buffer=%d recv %d\n", len(ch), data)
			if len(ch) == 0 {
				break LOOP
			}
		}
	}
}

// chanRecvClose 从已关闭的channel中接收数据，没有数据时会接收零值
func chanCloseRecv() {
	ch := make(chan int)
	close(ch)

	select {
	case data := <-ch:
		fmt.Println("recv ", data)
	}
}

func caseRecv() {
	//chanWaitRecv()
	chanBufferRecv()
	//chanCloseRecv()
}

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

//routineReturnErr 返回routine中的错误
func routineReturnErr() {
	type fun func() error
	tasks := []fun{
		func() error {
			return nil
		},
		func() error {
			return errors.New("wrong")
		},
	}

	errCh := make(chan error, len(tasks))
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))

	for i := range tasks {
		go func(i int) {
			defer wg.Done()

			if err := tasks[i](); err != nil {
				errCh <- err
			}
		}(i)
	}
	wg.Wait()

	select {
	case err := <-errCh:
		fmt.Println(err.Error())
	default:
		fmt.Println("normal")
	}
}

// func main() {
// 	caseSend()
// 	caseRecv()
// 	baseSelect()

// 	timeoutSelect()
// 	checkChanFullSelect()
// 	loopSelect()
// 	routineReturnErr()
// }
