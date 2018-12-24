package main

import (
	"fmt"
)

func test() int {
	a := 5

	return a
}

func main() {
	const PI = 3.1415926
	const a int = 6 / 3
	const b = 6

	const n = "hello"
	const l = len(n)

	//const b = test()
	//_ = b
	_ = l

	fmt.Println(PI)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(n)
	fmt.Println(l)
}
