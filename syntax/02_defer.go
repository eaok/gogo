package main

import (
	"fmt"
)

func f11() (result int) {
	result = 0
	defer func() {
		result++
	}()
	return
}

func f22() (r int) {
	t := 5
	r = t
	defer func() { //defer 函数被插入到赋值与返回之间执行，这个例子中返回值r没有被修改
		t = t + 5
	}()
	return
}

func f33() (t int) {
	t = 5
	defer func() {
		t = t + 5 //然后执行defer函数,t值被修改
	}()
	return
}

func f44() (r int) {
	r = 1
	defer func(r int) { //这里的r传值进去的，是原来r的copy，不会改变要返回的那个r值
		r = r + 5
	}(r)
	return
}

//deferCall 参数会被实时解析
func deferCall() {
	a, b := 1, 2
	defer calc("1", a, calc("10", a, b))

	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)

	return ret
}

//readReturnValue 可以读取有名返回值
func readReturnValue() (i int) {
	defer func() {
		i++
	}()

	return 1
}

func main() {
	// fmt.Println(f11())
	// fmt.Println(f22())
	// fmt.Println(f33())
	// fmt.Println(f44())
	deferCall()
	fmt.Println(readReturnValue())
}
