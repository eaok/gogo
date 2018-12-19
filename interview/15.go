//是否可以编译通过？如果通过，输出什么？
package main

import (
    "fmt"
)

func main() {
    list := new([]int)
//    list := make([]int, 0)
    list = append(list, 1)
    fmt.Println(list)
}
