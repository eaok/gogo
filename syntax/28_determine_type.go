package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "abc"
	i[2] = Student{"mike", 777}

	//comma-ok断言
	//value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int,内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string,内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为Student,内容name为%s,id=%d\n", index, value.name, value.id)
		}
	}

	//switch语句
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int,内容为%d\n", index, data)
		case string:
			fmt.Printf("x[%d] 类型为string,内容为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d] 类型为Student,内容name为%s,id=%d\n", index, value.name, value.id)
		}
	}
}
