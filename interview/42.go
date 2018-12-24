//请说出下面代码存在什么问题
package main

type student struct {
	Name string
}

func f(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		msg.Name
	}
}

//问题一：interface{}是一个没有声明任何方法的接口。
//问题二：Name是一个属性，而不是方法，interface{}类型的变量无法调用属性。
