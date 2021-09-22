package main

import (
	"context"
	"fmt"
	proto "micro_demo/logger/proto"

	"github.com/micro/go-micro/util/log"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "你好呀！" + req.Name
	return nil
}

// logWrapper1 包装HandlerFunc类型的接口
func logWrapper1(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Logf("[logWrapper1] %s 收到请求1", req.Endpoint())
		err := fn(ctx, req, rsp)
		log.Logf("[logWrapper1] %s 收到请求2", req.Endpoint())
		return err
	}
}

// logWrapper2 包装HandlerFunc类型的接口
func logWrapper2(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Logf("[logWrapper2] %s 收到请求1", req.Endpoint())
		err := fn(ctx, req, rsp)
		log.Logf("[logWrapper2] %s 收到请求2", req.Endpoint())
		return err
	}
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		// 声明包装器
		micro.WrapHandler(logWrapper1, logWrapper2),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
