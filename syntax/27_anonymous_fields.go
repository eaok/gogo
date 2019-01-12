package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	*Person //指针类型
	int     //基础类型的匿名字段
	string  //自定义类型

}

func main() {
	s := Student{&Person{"mike", 'm', 18}, 666, "go"}
	fmt.Println(s.name, s.sex, s.age, s.int, s.string)

	var s2 Student
	s2.Person = &Person{"go", 'm', 22}
	/*s2.Person = new(Person)
	s2.name = "go"
	s2.sex='m'
	s2.age = 18*/
	s2.int = 222
	s2.string = "sz"
	fmt.Println(s2.name, s2.sex, s2.age, s2.int, s2.string)
}
