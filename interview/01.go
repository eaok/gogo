//1、写出下面代码输出内容。
package main

import (
    "fmt"
)

func main()  {
    defer_call()
}

func defer_call()  {
    defer func() {
        fmt.Println("打印前")
    }()

    defer func() {
//        if err := recover(); err != nil {
//            fmt.Println(err) //err就是panic
//        }
        fmt.Println("打印中")
    }()

    defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

//panic 需要等defer 结束后才会向上传递。 出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。
