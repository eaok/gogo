//是否可以编译通过？如果通过，输出什么？
package main

import (
	"fmt"
    "reflect"
)

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
//    if reflect.DeepEqual(sm1, sm2) {
//        fmt.Println("sm1 == sm2")
//    } else {
//        fmt.Println("sm1 == sm2")
//    }
}

//结构体属性中有不可以比较的类型，如map,slice。不能比较
