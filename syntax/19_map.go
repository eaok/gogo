package main

import (
	"fmt"
	"sort"
)

type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var m map[string]string
	m = make(map[string]string, 10)
	m["a"] = "wangwu"
	m["b"] = "zhangsan"
	fmt.Println(m)

	m2 := make(map[string]string, 3)
	m2["a"] = "a"
	m2["b"] = "b"
	fmt.Println(m2)

	m3 := map[string]string{
		"n1": "x",
		"n2": "y",
		"n3": "z", //最后一个逗号不能去掉
	}
	m3["a"] = "a"
	fmt.Println(m3)

	//二维map
	var m4 map[string]map[string]string
	m4 = make(map[string]map[string]string, 2) //给分配大小
	m4["t1"] = make(map[string]string, 3)      //一维也是一个map
	m4["t1"]["n1"] = "t1n1"
	m4["t1"]["n2"] = "t1n2"

	m4["t2"] = make(map[string]string, 2)
	m4["t2"]["n1"] = "t2n1"
	m4["t2"]["n2"] = "t2n2"
	fmt.Println(m4)

	//map curd cu看上面
	delete(m4, "t1")
	//make(map[string]map[string]string) //清除所有的key
	fmt.Println(m4)
	val, ok := m3["t1"] //如果存在ok是 true,否则false, val 为对应的值
	if ok {
		fmt.Printf("t1存在 val为 %v\n", val)
	} else {
		fmt.Println("t1不存在")
	}

	//遍历,求长度也是用len(map)方法
	for i, v := range m4 {
		for j, val := range v {
			fmt.Printf("i=%v, j=%v, val=%v\n", i, j, val)
		}
	}

	//map 切片
	var mapSlice []map[string]string
	mapSlice = make([]map[string]string, 1)
	ms1 := map[string]string{
		"name": "zhangsan",
		"age":  "10",
	}
	mapSlice[0] = ms1
	ms2 := map[string]string{
		"name": "lisi",
		"age":  "20",
	}
	mapSlice = append(mapSlice, ms2)
	fmt.Println(mapSlice)

	//map排序ïkey放入切片,对切片排序;然后取值
	m5 := map[int]string{
		1: "av",
		2: "bv",
		3: "cv",
	}
	fmt.Println(m5)
	var keys []int
	for k, _ := range m5 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, v := range keys {
		fmt.Printf("key k=%v, v=%v\r\n", v, m5[v])
	}

	//map是引用传值
	test(m5)
	fmt.Println(m5) //[1:av 2:bv 3:233333]

	//map的value一般是用结构体
	m6 := make(map[string]Stu, 2)
	stu1 := Stu{"tom", 18, "北京"}
	stu2 := Stu{"王五", 28, "武汉"}
	m6["no1"] = stu1
	m6["no2"] = stu2

	for k, v := range m6 {
		fmt.Printf("学生号%v,name:%v,age:%v,address:%v\n", k, v.Name, v.Age, v.Address)
	}
}

func test(m map[int]string) {
	m[3] = "233333"
}
