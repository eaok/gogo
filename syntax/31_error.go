package main

import (
	"fmt"
)

func main() {
	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println(result)
	}
}

func MyDiv(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("%s", "分母不能为0")
		//err = errors.New("分母不能为0")
	} else {
		result = a / b
	}

	return
}
