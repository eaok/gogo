package main

import (
	"fmt"
	"io"
	"net"
	"protobuf/proto/message"

	"github.com/golang/protobuf/proto"
)

func main() {
	address := "localhost:8080"
	listener, err := net.Listen("tcp", address)
	if err != nil {

		fmt.Errorf("listen err:", err)

	}
	fmt.Println("[START] Server listenner: ", address)

	for {
		conn, err := listener.Accept()
		if err != nil {

			fmt.Println("conn err:", err)
			return

		}

		// 异步执行请求业务
		processing(conn)
	}
}

func processing(conn net.Conn) {
	// 延迟关闭
	defer conn.Close()
	// 缓冲
	buf := make([]byte, 4096)

	for {
		len, err := conn.Read(buf)
		// 读取结束
		if err == io.EOF {

			return

		}
		if err != nil {

			fmt.Println("conn read err:", err)
			return

		}
		user := &message.Message{}
		err = proto.Unmarshal(buf[:len], user)
		if err != nil {
			fmt.Println("proto unmarshal err:", err)
			return
		}

		fmt.Printf("receive data:%v  length:%v  ip:%v\n",
			user.Message, user.Length, conn.RemoteAddr())
	}
}
