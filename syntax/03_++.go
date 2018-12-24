package main

import "fmt"

func main() {
	n := 0
	p := &n

	n++
	fmt.Println(n)
	*p++
	fmt.Println(n)

	//b := n++
	//if n++ == 1 {}
	//++n
}
