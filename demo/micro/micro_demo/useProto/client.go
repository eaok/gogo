package main

import (
	"context"
	"fmt"
	proto "micro_demo/useProto/proto"

	micro "github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService()
	service.Init()

	//创建专属于Greeter服务的配套客户端
	greeter := proto.NewGreeterService("greeter.service", service.Client())

	//直接调用Greeter服务的Hello函数
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "wang"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
