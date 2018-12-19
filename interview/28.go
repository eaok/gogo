//编译执行下面代码会出现什么?
package main

func test() []func() {
    var funs []func()
    for i := 0; i < 2; i++ {
//		x := i
        funs = append(funs, func() {
//            println(&x, x)
            println(&i, i)
        })
    }

    return funs
}

func main() {
    funs := test()
    for _, f := range funs{
        f()
    }
}

//闭包延迟求值
//for循环复用局部变量i，每一次放入匿名函数的应用都是想一个变量。
