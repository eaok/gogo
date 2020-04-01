package main

import (
	"errors"
	"fmt"
)

// nameRetVar1() 闭包引用
/*
r = 0
defer
return
*/
func nameRetVar1() (r int) {
	defer func() {
		r++
	}()

	return 0 // 1
}

// nameRetVar2()
/*
r = t
defer
return
*/
func nameRetVar2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()

	return t // 5
}

// nameRetVar3() 函数调用
/*
r = 1
defer
return
*/
func nameRetVar3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)

	return 1 // 1
}

// nameRetNone1 闭包引用
func nameRetNone1() (r int) {
	r = 5
	defer func() {
		r = r + 5
	}()

	return // 10
}

// nameRetNone2 函数调用
func nameRetNone2() (r int) {
	r = 1
	defer func(r int) {
		r = r + 5
	}(r)

	return // 1
}

// anonymityRet1
/*
annoy := i
defer
return
*/
func anonymityRet1() int {
	var i int
	defer func() {
		i++
	}()

	return i // 0
}

// calc 返回两数的和
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)

	return ret
}

// nameRetVar 有名返回值函数return后有变量
func nameRetVar() {
	fmt.Println(nameRetVar1())
	fmt.Println(nameRetVar2())
	fmt.Println(nameRetVar3())
}

// nameRetNone 有名返回值函数return后没有变量
func nameRetNone() {
	fmt.Println(nameRetNone1())
	fmt.Println(nameRetNone2())
}

// anonymity 匿名返回值函数
func anonymityRet() {
	fmt.Println(anonymityRet1())
}

// deferCall 参数会被实时解析
func deferCall() {
	a, b := 1, 2
	defer calc("1", a, calc("10", a, b))

	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func main() {
	errors.New()
	nameRetVar()
	nameRetNone()
	anonymityRet()
	deferCall()
}
