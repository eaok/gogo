package main

import (
	"fmt"
	"net"
	"protobuf/proto/message"

	"github.com/gogo/protobuf/proto"
)

func main() {
	var msg string

	// 收集消息
	fmt.Println("请输入短消息:")
	if _, err := fmt.Scanf("%s", &msg); err != nil {
		fmt.Println("scanf error: ", err)
		return
	}

	// 编码数据
	user := &message.Message{
		Message: msg,
		Length:  *proto.Int(len(msg)),
	}
	pbData, err := proto.Marshal(user)
	if err != nil {
		fmt.Println("proto marshal err ", err)
		return
	}

	// 连接server
	address := "localhost:8080"
	conn, err := net.Dial("tcp", address)
	defer conn.Close()
	if err != nil {
		fmt.Println("net dial err ", err)
		return
	}

	// 发送数据
	n, err := conn.Write(pbData)
	if err != nil {
		fmt.Println("net write err ", err)
		return

	}
	fmt.Println("conn write len:", n)
}
