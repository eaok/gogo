package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type Something struct {
	Value []int
}

// 零切片
func zeroSlice() {
	var s = make([]int, 9)
	fmt.Println(s)
}

// 空切片
func emptySlice() {
	var s1 = []int{}
	var s2 = make([]int, 0)
	fmt.Println(*(*[3]int)(unsafe.Pointer(&s1)), len(s1), cap(s1), s1 == nil) // 第一个地址为所有空切片共享的地址
	fmt.Println(*(*[3]int)(unsafe.Pointer(&s2)), len(s2), cap(s2), s2 == nil)

}

// nil切片
func nilSlice() {
	var s3 []int
	var s4 = *new([]int)
	fmt.Println(*(*[3]int)(unsafe.Pointer(&s3)), len(s3), cap(s3), s3 == nil)
	fmt.Println(*(*[3]int)(unsafe.Pointer(&s4)), len(s4), cap(s4), s3 == nil)

}

func nilBetterReason() {
	var emptySlice []int
	var nilSlice = []int{}

	fmt.Printf("emptySlice %v %#v\n", emptySlice == nil, emptySlice) // emptySlice true []int(nil)
	fmt.Printf("nilSlice %v %#v\n", nilSlice == nil, nilSlice)       // nilSlice false []int{}

	// 空切片和nil切片有时候会隐藏在结构体中，导致结果不一样
	var emptySlice1 = Something{}
	var nilSlice1 = Something{[]int{}}
	fmt.Println(emptySlice1.Value == nil, nilSlice1.Value == nil) // true false

	// 还有一个极为不同的地方在于JSON序列化
	var emptySlice2 = Something{}
	var nilSlice2 = Something{[]int{}}
	emptySlicebs, _ := json.Marshal(emptySlice2)
	nilSlicebs, _ := json.Marshal(nilSlice2)
	fmt.Printf("emptySlicebs %v\n", string(emptySlicebs)) // emptySlicebs {"Value":null}
	fmt.Printf("nilSlicebs %v\n", string(nilSlicebs))     // nilSlicebs {"Value":[]}
}

// func main() {
// 	zeroSlice()
// 	emptySlice()
// 	nilSlice()

// 	// 空切片和nil切片使用上有区别，官方推荐使用nil切片
// 	nilBetterReason()
// }
