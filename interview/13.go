//下面函数有什么问题？
package main

import (
    "fmt"
)

func main() {
    a, err := funcMui(1, 2)
    fmt.Println(a)
    _ = err
}

func funcMui(x, y int) (sum int, error){
//func funcMui(x, y int) (int, error){
    return x + y, nil
}
