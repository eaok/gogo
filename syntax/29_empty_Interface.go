package main

import "fmt"

func xxx(arg ...interface{}) {

}

func main() {
	//空接口是万能类型，保存任意类型的值
	var i interface{} = 1
	fmt.Println("i=", i)

	i = "abc"
	fmt.Println("i=", i)
	xxx("aaaa", 5)
}
