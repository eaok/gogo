//执行下面的代码发生什么？
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				//                close(ch)
				return
			}
			fmt.Println("a: ", a)
		}
	}()

	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
}

//往已经关闭的channel写入数据会panic的。
