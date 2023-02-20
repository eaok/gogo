package main

import (
	"strconv"
	"sync"
	"time"

	cmap "github.com/orcaman/concurrent-map"
)

// map是线程不安全的,即使并发读写没有冲突也会报错(fatal error: concurrent map read and map write)
func mapSecurity() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	m := make(map[int]int)

	go func() {
		for i := 0; i < 9999; i++ {
			if j, ok := m[i]; ok {
				println(j)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 9999; i++ {
			m[i] = i * 10
		}
		wg.Done()
	}()

	wg.Wait()
}

// 使用读写锁sync.RWMutex/互斥锁sync.Mutex解决map并发安全问题
func mapUseSyncMutex() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	m := struct {
		set map[int]int
		sync.Mutex
	}{set: make(map[int]int)}

	go func() {
		for i := 0; i < 9999; i++ {
			m.Lock()
			if j, ok := m.set[i]; ok {
				println(j)
			}
			m.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 9999; i++ {
			m.Lock()
			m.set[i] = i * 10
			m.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
}

// 使用sync.Map解决map并发安全问题
func mapUseSyncMap() {
	m := sync.Map{}
	go func() {
		for {
			_, _ = m.Load(1)
		}
	}()
	go func() {
		for {
			m.Store(1, 2)
		}
	}()

	time.Sleep(time.Second)
}

// 标准库中的sync.Map是专为append-only场景设计的，如果想将Map用于一个类似内存数据库，可以用concurrent-map
func mapUseConcurrent() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	m := cmap.New()

	go func() {
		for i := 0; i < 9999; i++ {
			if j, ok := m.Get(strconv.Itoa(i)); ok {
				println(j)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 9999; i++ {
			m.Set(strconv.Itoa(i), i*10)
		}
		wg.Done()
	}()

	wg.Wait()
}

// 对于不容易发现的并发问题,可以使用-race参数进行并发检测
// go run --race summary/12_mapSecuritySum.go
// func main() {
// 	mapSecurity()
// 	mapUseSyncMutex()
// 	mapUseSyncMap()
// 	mapUseConcurrent()
// }
