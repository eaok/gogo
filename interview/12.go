//是否可以编译通过？如果通过，输出什么？
package main

func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

func GetValue() int {
	return 1
}

//编译失败，因为type只能使用在interface
