package main

import (
	"fmt"
)

func main() {
	var intArr [5]int = [...]int{1, 2, 3, 4, 5} //定义一个数组
	slice := intArr[2:4]                        //第二个(包含)下标到第四个下标(不包含)

	fmt.Println("slice = ", slice) //[3 4]
	fmt.Println("slice len = ", len(slice))
	fmt.Println("slice cap = ", cap(slice)) //切片的容量

	//通过make创建切片
	var makeSlice []int = make([]int, 2, 5) //第一个参数为类型, 第二个参数为len, 第三个为cap
	fmt.Println("makeSlice = ", makeSlice)  //数字类型(int, float)默认为0, sting 为'', bool为false  和数组一样
	fmt.Println("makeSlice len = ", len(makeSlice))
	fmt.Println("makeSlice cap = ", cap(makeSlice))

	//类似于make
	var mslice []string = []string{"zhangsan", "lisi", "wangwu", ""}
	mslice[3] = "ermazi"
	fmt.Println("mslice = ", mslice)
	fmt.Println("mslice len = ", len(mslice))
	fmt.Println("mslice cap = ", cap(mslice))

	//遍历
	for i := 0; i < len(mslice); i++ {
		fmt.Printf("i = %v, v = %v\n", i, mslice[i])
	}

	for j, mv := range mslice {
		fmt.Printf("j = %v, mv = %v\n", j, mv)
	}

	//细节
	var testArr [5]int = [...]int{1, 2, 3, 4, 5}
	testSlice := testArr[0:len(testArr)] //取出数组里面所有值 testArr[:]可以简写成这样
	startSlice := testArr[:4]            //从0开始到第四个下标
	endSlice := testArr[2:]              // 从2开始到结束

	testSlice2 := startSlice[1:3] //slice也可以从slice种切,和数组的规则一样
	testSlice[2] = 100            //slice 是引用，更改了值都会影响其它关联的指向变量的值(testArr, starSlice对应的都改变了)

	fmt.Println("testArr = ", testArr)
	fmt.Println("testSlice = ", testSlice)
	fmt.Println("testSlice2 = ", testSlice2)
	fmt.Println("endSlice = ", endSlice) //testArr的值改变了;它也得改变

	//append,copy的使用
	var apSlice []int = make([]int, 5, 10)
	apSlice = append(apSlice, 1, 2, 3, 4, 5) //追加元素
	fmt.Println("apSlice = ", apSlice)       //[0 0 0 0 0 1 2 3 4 5]

	apSlice = append(apSlice, apSlice...) //也可以追加一个slice
	fmt.Println("apSlice = ", apSlice)    // [0 0 0 0 0 1 2 3 4 5 0 0 0 0 0 1 2 3 4 5]

	var coppySlice []string = []string{"zhangsan", "lisi", "wangwu"}
	var dslice []string = make([]string, 3, 3)
	copy(dslice, coppySlice)         //把copySlice的值拷贝到dslice种
	fmt.Println("dslice = ", dslice) //[zhangsan lisi wangwu]

	var minSlice []string = make([]string, 1, 1) //只有一个长度
	copy(minSlice, coppySlice)                   //此时不会出现越界的错误,只是拷贝了一个值
	fmt.Println("minSlice = ", minSlice)         //[zhangsan]

	//字符串底层是一个byte数组,所以可以通过slice来操作
	str := "hello,gogogo"
	strSlice := str[2:6] //这个返回的string,感觉应该是slice
	fmt.Printf("strSlice = %v, %T\n", strSlice, strSlice)

	arr := []byte(str) //string 转换成byte切片
	fmt.Printf("arr = %v, %T\n", arr, arr)
	arr[0] = 'H'
	//arr[1] = '我' //中文会报错
	str = string(arr[:]) //[]byte切片转换成字符串
	fmt.Println(str)     //Hello,gogogo

	//通过[]rune切片解决中文
	arr2 := []rune(str)
	arr2[1] = '我'
	str = string(arr2)
	fmt.Println(str) //H我llo,gogogo

	//练习 斐波那契
	fmt.Println(test(8)) //[1 1 2 3 5 8 13 21]
}

func test(n int) []uint64 {
	var slice []uint64 = make([]uint64, n)
	for i := 0; i < n; i++ {
		if i <= 1 {
			slice[i] = 1
		} else {
			num := slice[i-1] + slice[i-2]
			slice[i] = num
		}
	}
	return slice
}
