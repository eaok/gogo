//编译执行下面代码会出现什么?
package main

func main() {

	for i := 0; i < 10; i++ {
	loop:
		println(i)
	}

	goto loop
}

//goto不能跳转到其他函数或者内层代码
