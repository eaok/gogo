package main

import (
    "fmt"
)

func main() {
    s1 := "hello"
    fmt.Println(s1[0])
//    fmt.Println(&s[1])
//标准索引法来获取字符串中的内容只对ASCII码字符串有效；另外，尝试获取字符串中某个字节的地址是非法的

    s2 := "Beginning of the string " + //自动添加分号,所以要写还这行末尾
        "second part of the string"
    fmt.Println(s2)
}
