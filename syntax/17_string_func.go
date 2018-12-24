package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello老美"                     //中文占三个字节
	fmt.Println("str的长度为len=", len(str)) //长度为11

	for i := 0; i < len(str); i++ {
		fmt.Printf("字符=%c\n", str[i]) //中文会乱码
	}

	rstr := []rune(str)                    //变成一个rune切片就可以解决中文
	fmt.Println("rstr的长度为len=", len(rstr)) //长度为7
	for i := 0; i < len(rstr); i++ {
		fmt.Printf("字符=%c\n", rstr[i]) //中文正常显示
	}

	//字符串转整数
	nstr, err := strconv.Atoi("123")
	fmt.Printf("nstr=%v,err=%v\n", nstr, err)

	//可以检验一个字符串是否为数字
	nstr2, err2 := strconv.Atoi("abc")
	if nil != err2 {
		fmt.Printf("转换失败 %v\n", err2)
	} else {
		fmt.Printf("转换成功 %v\n", nstr2)
	}

	//整数转字符串
	atostr := strconv.Itoa(12345)
	fmt.Printf("nstr=%v,err=%T\n", atostr, atostr) //12345 string

	//字符串转[]byte(string) ascii码
	by := []byte("hello,abc")
	fmt.Println("byte=", by) //每个字符的ascii码

	//[]byte{11,12,13} 转字符串 string([]byte{11,12})
	bystr := string([]byte{104, 101, 108, 108, 111, 44, 97, 98, 99})
	fmt.Println("bystr=", bystr) //hello,abc

	//十进制转换成指定进制(2,8,16)
	num2 := strconv.FormatInt(123, 2) //转成2进制
	fmt.Println("num2=", num2)

	num8 := strconv.FormatInt(123, 8) //转成8进制
	fmt.Println("num2=", num8)

	//判断一个字符串是否在另外一个字符串中出现
	b := strings.Contains("abczxxxxxx", "y") //返回bool型
	fmt.Println("b=", b)

	//一个字符串在另外一个字符串中出现得次数
	ct := strings.Count("wahaha", "ha")
	fmt.Println("ct=", ct)

	//判断两个字符串是否相等 == 区分大小写,使用strings.EqualFold()不区分
	se := "abc" == "Abc"
	sen := strings.EqualFold("abc", "Abc")
	fmt.Println("se=", se)   //false
	fmt.Println("sen=", sen) //true

	//字符串在某一字符串中首次出现得位置 -1 是不存在
	index := strings.Index("wahaha", "h")   //2
	index_1 := strings.Index("wahaha", "z") //-1
	fmt.Println("index=", index)            //false
	fmt.Println("index_1=", index_1)        //true

	//最后出现得位置 -1不存在
	lastIndex := strings.LastIndex("wahaha", "ha")
	fmt.Println("lastIndex=", lastIndex) //4

	//字符串替换
	originstr := "wahaha"
	newstr := strings.Replace(originstr, "ha", "wa", 1) // 1是替换一个,2替换两个, -1 替换所有
	fmt.Println("newstr=", newstr)                      //wawaha

	//把一个字符串切割成数组
	arrstr := "ni,hen,niu,b"
	arr := strings.Split(arrstr, ",")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Printf("arr=%v,类型%T\n", arr, arr) //wawaha

	//字符串转换大小写
	fmt.Println(strings.ToUpper("abc")) //ABC
	fmt.Println(strings.ToLower("AbC")) //abc

	//去掉字符串左右空格
	fmt.Printf("%q\n", strings.TrimSpace("  xyz   ")) //"xyz" %q 加上了""

	//指定去掉左右两边的某个字符 TrimLeft,TrimRight 去掉左或者右边
	fmt.Println(strings.Trim("!!wahah!ahahah!!", "!")) //wahah!ahahah 中间的不会去掉

	//是否以某个字符串开头(HasPrefix);或者结尾(HasSuffix)
	fmt.Println(strings.HasPrefix("http://www.wahaha.com", "http"))         //true
	fmt.Println(strings.HasSuffix("http://www.wahaha.com/test.avi", "avi")) //true
}
