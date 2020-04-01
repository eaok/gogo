package main

import (
	"context"
	"fmt"
	"log"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

func main() {
	service := micro.NewService()
	service.Init()
	c := service.Client()

	//客户端给Node发送Grpc请求
	req := "wang"
	request := c.NewRequest("service.greeter", "Greeter.Hello", req,
		client.WithContentType("application/json"))

	var response string
	if err := c.Call(context.TODO(), request, &response); err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}
