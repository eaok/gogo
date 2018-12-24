//下面的迭代会有什么问题？
package main

import (
	"fmt"
	"sync"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	//ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}

		close(ch)
		set.RUnlock()
	}()
	return ch
}

func main() {
	th := threadSafeSet{
		s: []interface{}{"1", "2", "3", "4", "5", "6", "7", "8"},
	}
	v := <-th.Iter()
	fmt.Sprintf("%s%v", "ch", v)
}
