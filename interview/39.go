//ABCD中哪一行存在错误？
package main

type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D
}

//函数中func f(x interface{})的interface{}可以支持传入golang的任何类型，包括指针，但是函数func g(x *interface{})只能接受*interface{}
