package main

import (
	"flag"
	"fmt"
	"rpc/grpc/withEtcd/etcdservice"
	pb "rpc/grpc/withEtcd/protoes"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var (
	ServiceName = flag.String("ServiceName", "hello_service", "service name")
	EtcdAddr    = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()
	r := etcdservice.NewResolver(*EtcdAddr)
	resolver.Register(r)

	// 客户端连接服务器
	// conn, err := grpc.Dial("127.0.0.1:3000", grpc.WithInsecure())
	conn, err := grpc.Dial(r.Scheme()+"://author/"+*ServiceName,
		grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败", err)
	}
	defer conn.Close()

	// 获得grpc句柄
	c := pb.NewHelloServerClient(conn)
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		// 远程单调用 SayHi 接口
		r1, err := c.SayHi(
			context.Background(),
			&pb.HelloRequest{
				Name: "Kitty",
			},
		)
		if err != nil {
			fmt.Println("Can not get SayHi:", err)
			return
		}
		fmt.Printf("%v: SayHi 响应：%s\n", t, r1.GetMessage())

		// 远程单调用 GetMsg 接口
		r2, err := c.GetMsg(
			context.Background(),
			&pb.HelloRequest{
				Name: "Kitty",
			},
		)
		if err != nil {
			fmt.Println("Can not get GetMsg:", err)
			return
		}
		fmt.Printf("%v: GetMsg 响应：%s\n", t, r2.GetMsg())
	}
}
