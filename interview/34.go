//执行下面的代码发生什么？
package main

import "fmt"

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

func main() {
	c := &ConfigOne{}
	c.String()
}

//如果类型实现String()，％v和％v格式将使用String()的值。因此，对该类型的String()函数内的类型使用％v会导致无限递归。
