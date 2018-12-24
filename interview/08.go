//下面的代码有什么问题?
package main

import (
	"fmt"
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	a := UserAges{make(map[string]int)}

	go func() {
		for i := 0; i < 1000; i++ {
			a.Add("TEST", 12)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(a.Get("TEST"))
		}
	}()

	time.Sleep(time.Hour)
}
