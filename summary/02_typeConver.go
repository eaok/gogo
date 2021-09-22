package main

import (
	"fmt"
	"strconv"
)

func stringToInt() {
	str := "9"

	// string to int
	num1, err := strconv.Atoi(str) //第一种
	//num1, err := strconv.ParseInt(str, 10, 0)	//第二种
	if err != nil {
		panic(err)
	}
	fmt.Printf("stringToInt %v %T\n", num1, num1)

	// string to int64
	num2, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("stringToInt %v %T\n", num2, num2)
}

func intToString() {
	var num1 int = 99
	var num2 int64 = 999

	// int to string
	str1 := strconv.Itoa(num1) //第一种
	//str1 := fmt.Sprintf("%d", num1)				//第二种
	//str1 := strconv.FormatInt(int64(num1), 10)	//第三种
	fmt.Printf("intToString %v %T\n", str1, str1)

	// int64 to string
	str2 := strconv.FormatInt(num2, 10)
	fmt.Printf("intToString %v %T\n", str2, str2)
}

func stringToFloat() {
	str := "3.1415926535"
	f1, _ := strconv.ParseFloat(str, 64)
	fmt.Printf("stringToFloat\t%v\t%T\n", f1, f1)
}

func floatToString() {
	f := 3.1415926535
	str := strconv.FormatFloat(f, 'f', -1, 64)
	fmt.Printf("floatToString\t%v\t%T\n", str, str)
}

func intToFloat() {
	num := 9999

	f1 := float32(num)
	f2 := float64(num)
	fmt.Printf("intToFloat\t%v\t%T\n", f1, f1)
	fmt.Printf("intToFloat\t%v\t%T\n", f2, f2)
}

func floatToInt() {
	f := 3.1415926535

	num := int(f)
	fmt.Printf("floatToInt\t%v\t%T\n", num, num)
}

func main() {
	stringToInt()
	intToString()
	stringToFloat()
	floatToString()
	intToFloat()
	floatToInt()
}
