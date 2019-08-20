package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//接收文件
func RecvFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err=", err)
		return
	}
	buf := make([]byte, 1024*4)

	//接收多少，写多少
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err=", err)
			}
			return
		}

		f.Write(buf[:n])
	}
}

func main() {
	//监听
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}
	defer listenner.Close()

	//阻塞用户连接
	conn, err1 := listenner.Accept()
	if err1 != nil {
		fmt.Println("listenner.Accept err1 =", err1)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err =", err)
		return
	}

	fileName := string(buf[:n])

	//回复“ok”
	conn.Write([]byte("ok"))
	//调用保存文件的方法
	RecvFile(fileName, conn)
}
