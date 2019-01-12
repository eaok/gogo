package main

import "fmt"

type Humaner interface {
	sayHi()
}

type Student struct {
	name string
	id   int
}

type Teacher struct {
	addr  string
	group string
}

type Mystr string

func (tmp *Student) sayHi() {
	fmt.Printf("Student[%s,%d] sayhi\n", tmp.name, tmp.id)
}

func (tmp *Teacher) sayHi() {
	fmt.Printf("Teacher[%s,%s] sayhi\n", tmp.addr, tmp.group)
}

func (tmp *Mystr) sayHi() {
	fmt.Printf("mystr[%s] sayhi\n", *tmp)
}

//只有一个函数，可以有不同表现，多态
func WhoSayHi(i Humaner) {
	i.sayHi()
}

func main() {
	//定义接口类型变量
	var i Humaner

	//如果定义的类型实现了某个接口类型声明的一组方法，那么这个类型的值就可以赋给这个接口类型
	s := &Student{"mike", 666}
	i = s
	i.sayHi()

	t := &Teacher{"bj", "go"}
	i = t
	i.sayHi()

	var str Mystr = "hello world"
	i = &str
	i.sayHi()
	fmt.Printf("-----------------\n")

	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str
	for _, i := range x {
		i.sayHi()
	}
	fmt.Printf("-----------------\n")

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)
}
