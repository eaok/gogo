package main

type Tester interface {
	Echo(string)
}

type Test1 struct {
	name string
}
type Test2 struct {
	name string
}

func (t *Test1) Echo(str string) {
	println("this is test1", str, t.name)
}

func (t *Test2) Echo(str string) {
	println("this is test2", str, t.name)
}

var test1 = []Tester{
	&Test1{},
	&Test2{},
}

func main() {
	test1[0] = &Test1{"wang"}

	for _, fun := range test1 {
		fun.Echo("hello")
	}
}
