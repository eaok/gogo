//下面代码能运行吗？为什么
package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {

	// 创建Show结构体对象
	s := new(Show)
	// 为字典Param赋初始值
	//    s.Param = Param{}
	// 修改键值对
	s.Param["RMB"] = 10000
	fmt.Println(s)
}

//字典Param的默认值为nil，当给字典nil增加键值对是就会发生运行时错误
