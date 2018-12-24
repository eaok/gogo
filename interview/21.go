//编译执行下面代码会出现什么?
package main

var (
    size := 1024
//	size     = 1024
	max_size = size * 2
)

func main() {
	println(size, max_size)
}
