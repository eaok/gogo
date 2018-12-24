//请说明下面代码书写是否正确
package main

import (
	"fmt"
	"sync/atomic"
)

var value int32

func SetValue(delta int32) {
	for {
		v := value
		// 比较并交换
		if atomic.CompareAndSwapInt32(&value, v(v+delta)) {
			//        if atomic.CompareAndSwapInt32(&value, v,v+delta) {
			fmt.Println(value)
			break
		}
	}
}

func main() {
	SetValue(100)
}
