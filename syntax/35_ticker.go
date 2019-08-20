package main

import (
	"fmt"
    "time"
    "runtime"
)

func main() {
    runtime.GOMAXPROCS(2)

	//创建定时器，每隔 1 秒后，定时器就会给 channel 发送一个事件(当前时间)
	ticker := time.NewTicker(time.Second * 1)
	i := 0

	go func() {
		for {
			<-ticker.C
			i++
			fmt.Println("i = ", i)
			if i == 5 {
				ticker.Stop() //停止定时器
                break
			}
		}
	}()

	for {
	}
}
