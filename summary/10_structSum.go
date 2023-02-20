package main

import (
	"fmt"
	"unsafe"
)

//structlayout -json . ST|structlayout-svg -t "align-guarantee" > 10_ag.svg
type ST struct {
	b   bool
	num int
	str string      // 16
	ptr *string     // 8
	arr [2]int      // 2 * 8
	sli []int       // 24
	i   interface{} // 16
	m   map[string]int
	st1 struct{}             // 0
	st2 struct{ str string } //16
	st3 struct{}             // 1
}

// structlayout -json . tooMuchPadding|structlayout-optimize -r
type tooMuchPadding struct {
	i16 int16
	i64 int64
	i8  int8
	i32 int32
	ptr *string
	b   bool
}

// zeroSizeField
func zeroSizeField() {
	type T1 struct {
		a struct{}
		x int64
	}

	type T2 struct {
		x int64
		a struct{}
	}
	a1 := T1{}
	a2 := T2{}
	fmt.Printf("zero size struct{} in field:\n"+
		"T1 (not as final field) size: %d\n"+
		"T2 (as final field) size: %d\n",
		unsafe.Sizeof(a1), // 8
		unsafe.Sizeof(a2)) // 64位：16；32位：12
}

// func main() {
// 	zeroSizeField()
// }
