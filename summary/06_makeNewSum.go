package main

import "fmt"

// makeExample make用于创建切片、哈希表和 Channel
func makeExample() {
	slice := make([]int, 0, 100)
	hash := make(map[int]bool, 10)
	ch := make(chan int, 5)

	fmt.Println(slice, hash, ch)
}

// newExample new创建一个类型并返回指向该类型的指针
func newExample() {
	i := new(int)
	//等价于
	//var v int
	//i := &v

	fmt.Println(*i)
}

func main() {
	makeExample()
	newExample()
}
