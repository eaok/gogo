package main

import (
	"fmt"
)

func main() {
	a := []int{
		1, 2, 3,
		4, 5, 6,
	}

	b := []int{
		1, 2, 3,
		4, 5, 6}
	c := []int{1, 2, 3, 4, 5, 6}

	d := []int{
		1, 2, 3,
		//4, 5, 6 //要有逗号
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
