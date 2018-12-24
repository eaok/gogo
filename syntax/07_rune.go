package main

import (
	"fmt"
)

func main() {
	s := "Go编程"
	fmt.Print(len(s))
	fmt.Println(len(s))
	fmt.Println(len(string(rune('编'))))
	fmt.Println(len([]rune(s)))
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2])
	fmt.Println(s[3])
}
