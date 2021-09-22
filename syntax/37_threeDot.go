package main

import (
	"fmt"
)

//1.用于接收不确定数量的参数
func show(args ...int) {
	for _, v := range args {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6}
	array := [...]int{7, 8} //同1

	//2.将slice打散进行传递
	slice1 = append(slice1, slice2...)
	show(slice1...)
	fmt.Println(array)
}

//3. golang里./...表示所有子目录
// go test ./...
