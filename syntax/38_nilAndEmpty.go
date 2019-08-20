package main

import "fmt"

//testSlice nil切片与空切片
func testSlice() {
	var slice1 []int //nil切片
	slice2 := make([]int, 0)
	slice3 := []int{}

	fmt.Printf("slice1: %T %t\n", slice1, slice1 == nil)
	fmt.Printf("slice2: %T %t\n", slice2, slice2 == nil)
	fmt.Printf("slice3: %T %t\n", slice3, slice3 == nil)
}

//testMap nil映射与空映射
func testMap() {
	var map1 map[string]string //nil映射不能存储键值对，仅起声明变量的作用
	map2 := map[string]string{}
	map3 := make(map[string]string)

	fmt.Printf("map1 %T %t\n", map1, map1 == nil)
	fmt.Printf("map2 %T %t\n", map2, map2 == nil)
	fmt.Printf("map3 %T %t\n", map3, map3 == nil)
}

func main() {
	testSlice()
	testMap()
}
