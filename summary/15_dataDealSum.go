package main

import (
	"fmt"
	"time"
)

// 格式化为年月日
func formatTime() {
	today := time.Now().Format("2006-01-02")
	fmt.Println(today) //2022-07-22

	//格式化为年月日时分秒
	datetime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(datetime) //2022-07-22 11:56:31
}

// 字符串转日期
func stringToDate() {
	str := "2017-09-13 21:57:01"
	loc, _ := time.LoadLocation("Local")
	timeval, _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	fmt.Println(timeval)
}

// 获取年月日
func getDateProperty() {
	day := time.Now().Day()
	month := time.Now().Format("01")
	year := time.Now().Year()
	fmt.Println(year, month, day)
}

// 获取时间戳
func getDateTimestamp() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixMilli()) // UnixMilli 和 UnixMicro 是go1.17版后新加的
	fmt.Println(time.Now().UnixMicro())
	fmt.Println(time.Now().UnixNano())
}

// func main() {
// 	// formatTime()
// 	stringToDate()
// 	getDateProperty()
// 	getDateTimestamp()
// }
