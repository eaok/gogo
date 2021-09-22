package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type teacher struct {
	Name string
	Age  int
	Sex  int
}

//shallowCopy go默认的是浅拷贝
func shallowCopy() {
	teacher1 := &teacher{
		Name: "zhao",
		Age:  24,
		Sex:  1,
	}

	teacher2 := teacher1
	fmt.Printf("%T, %v, %p \n", teacher1, teacher1, teacher1)
	fmt.Printf("%T, %v, %p \n", teacher2, teacher2, teacher2)

	teacher2.Name = "qian"
	fmt.Printf("%T, %v, %p \n", teacher1, teacher1, teacher1)
	fmt.Printf("%T, %v, %p \n", teacher2, teacher2, teacher2)
}

// deepCopy 基于序列化和反序列化来实现对象的深度拷贝
// 需要深拷贝的变量必须首字母大写才可以被拷贝
func deepCopyStruct(dst, src interface{}) error {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// deepCopyStructTest 测试结构体的深拷贝
func deepCopyStructTest() {
	teacher1 := &teacher{
		Name: "zhao",
		Age:  24,
		Sex:  1,
	}

	teacher2 := new(teacher)
	if err := deepCopyStruct(teacher2, teacher1); err != nil {
		panic(err.Error())
	}
	fmt.Printf("%T, %v, %p \n", teacher1, teacher1, teacher1)
	fmt.Printf("%T, %v, %p \n", teacher2, teacher2, teacher2)
	teacher2.Name = "qian"
	fmt.Printf("%T, %v, %p \n", teacher1, teacher1, teacher1)
	fmt.Printf("%T, %v, %p \n", teacher2, teacher2, teacher2)
}

//deepCopySlice 切片可以使用copy来实现深拷贝
func deepCopySlice() {
	s := []int{1, 2, 3}
	sdst := make([]int, 3)

	copy(sdst, s)
	fmt.Printf("%T, %v, %p \n", s, s, s)
	fmt.Printf("%T, %v, %p \n", sdst, sdst, sdst)
}

//deepCopyTest 深拷贝测试
func deepCopyTest() {
	deepCopyStructTest()
	deepCopySlice()
}

func main() {
	shallowCopy()
	deepCopyTest()
}
