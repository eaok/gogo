package main

import "time"

// classicLoop 经典循环
func classicLoop() {
	flag := true

	// loop1
	for i := 0; i < 5; i++ {
		println("classic loop 1")
	}

	// loop2
	for flag {
		println("classic loop2")
	}

	// loop3
	for {
		println("classic loop3")
	}
}

// forArraySlice 遍历数组切片
func forArraySlice() {
	data := []int{1, 2, 3}

	// 不关心索引和数据
	for range data {
		println("range loop1 for slice")
	}

	// 只关心索引
	for i := range data {
		println("range loop2 for slice ", i)
	}

	// 关心索引和数据
	for i, elem := range data {
		println("range loop3 for slice ", i, elem)
	}
}

// forString 遍历字符串
func forString() {
	data := "hello"

	// 和数组切片类似，字符串底层会转成rune切片，然后再遍历
	for i, r := range data {
		println("range loop for string ", i, r)
	}
}

// forMap 哈希表的遍历
func forMap() {
	data := make(map[string]int, 3)
	data["zhao"] = 1
	data["qian"] = 3

	for range data {
		println("range loop for map")
	}

	for k := range data {
		println("range loop for map ", k)
	}

	for k, v := range data {
		println("range loop for map ", k, v)
	}
}

// forChan 遍历channel
func forChan() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(in chan int) {
		for i := 0; i < 3; i++ {
			in <- i
		}

		close(in)
	}(ch1)

	go func(in chan int) {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			in <- i
		}

		close(in)
	}(ch2)

	for range ch1 {
		println("range loop for chan")
	}

	for v := range ch2 {
		println("range loop for chan ", v)
	}
}

// rangeLoop 范围循环
func rangeLoop() {
	//forArraySlice()
	//forString()
	//forMap()
	forChan()
}

func main() {
	//classicLoop()
	rangeLoop()
}
