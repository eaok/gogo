package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaa")
}

func testb(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	//显示调用panic，导致程序中断
	panic("this is a panic test")
	//var a [10]int
	//a[x] = 111 //当x为22时候，导致数组越界，产生一个panic
}

func testc() {
	fmt.Println("cccccc")
}

func main() {
	testa()
	testb(20)
	testc()
}
