package main

import (
	"fmt"
	"reflect"
)

type I1 interface{ fun() }
type I2 interface{ fun() }
type S1 struct{ name string }
type S2 struct{ name string }

func (s S1) fun() {}
func (s S2) fun() {}

func compareInterface() {
	var a, b, c, d I1
	var e I2
	a = S1{"hello"}
	b = S1{"hello"}
	c = S1{"world"}
	d = S2{"hello"}
	fmt.Println(a == b) //true
	fmt.Println(a == c) //false
	fmt.Println(a == d) //false
	fmt.Println(a == e) //false
}

func compareSlice() {
	a1 := []int{1, 2}
	a2 := []int{1, 2}
	if reflect.DeepEqual(a1, a2) {
		fmt.Println(a1, "==", a2)
	}
}

type S struct{ s string }

func compareStruct() {
	s1 := S{s: "hello"}
	s2 := S{s: "hello"}
	if reflect.DeepEqual(s1, s2) {
		fmt.Println(s1, "==", s2)
	}
}

func compareMap() {
	m1 := map[int]string{1: "a", 2: "b"}
	m2 := map[int]string{1: "a", 2: "b"}
	if reflect.DeepEqual(m1, m2) {
		fmt.Println(m1, "==", m2)
	}
}

func main() {
	compareInterface()
	compareSlice()
	compareStruct()
	compareMap()
}
