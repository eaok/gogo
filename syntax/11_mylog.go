// mylog project main.go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	log.Println("func init()")                           //默认只有日期和时间
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime) //加上文件名称
}

func main() {
	errFile, err := os.OpenFile("errFile.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0655)
	if err != nil {
		log.Fatalln("Open file errFile.log error: ", err)
	}

	Info = log.New(os.Stdout, "[Info]", log.Llongfile|log.LstdFlags) // 日期时间+完整文件路径名称
	Warning = log.New(os.Stdout, "[Warning]", log.Llongfile|log.LstdFlags)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "[Error]", log.Llongfile|log.LstdFlags) //同时输出到标准错误和文件中

	Info.Println("这是一条 Info 日志信息!")
	Warning.Println("这是一条 Warning 日志信息")
	Error.Println(" 这是一条 Wrror 日志信息， 这条信息还会输出到 errFile.log 文件中去!")

	log.Println("func main()")
	log.Println("Hello World!")

	fmt.Println("Hello World!")
}
