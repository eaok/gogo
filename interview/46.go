//请说出下面代码哪里写错了。
package main

import (
	"fmt"
	"time"
)

func main() {
	abc := make(chan int, 1000)
	for i := 0; i < 10; i++ {
		abc <- i
	}

	go func() {
		for {
			a := <-abc
			//            a,ok := <-abc
			//            if !ok {
			//                fmt.Println("结束！")
			//                return
			//            }
			fmt.Println("a: ", a)
		}
	}()

	close(abc)
	fmt.Println("close")
	time.Sleep(time.Second * 10)
}

//go中的for循环是死循环，应该设置出口。
