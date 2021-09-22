package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"rpc/grpc/withEtcd/etcdservice"
	pb "rpc/grpc/withEtcd/protoes"
	"syscall"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
)

var host = "127.0.0.1"
var (
	ServiceName = flag.String("ServiceName", "hello_service", "service name")
	Port        = flag.Int("Port", 3000, "listening port")
	EtcdAddr    = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address")
)

// 对象要和proto内定义的服务一致
type server struct {
}

// 实现rpc的 SayHi接口
func (s *server) SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReplay, error) {
	fmt.Printf("SayHi: %s, %s\n", in.Name, time.Now().Format(TIME_FORMAT))
	return &pb.HelloReplay{
		Message: "Hi " + in.Name,
	}, nil
}

// 实现rpc的 GietMsg接口
func (s *server) GetMsg(ctx context.Context, in *pb.HelloRequest) (*pb.HelloMessage, error) {
	fmt.Printf("GetMsg: %s, %s\n", in.Name, time.Now().Format(TIME_FORMAT))
	return &pb.HelloMessage{
		Msg: "Server msg is coming...",
	}, nil
}

func main() {
	flag.Parse()
	// 监听网络
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *Port))
	if err != nil {
		fmt.Println("网络异常：", err)
		return
	}
	defer ln.Close()

	// 创建grpc句柄
	srv := grpc.NewServer()
	defer srv.GracefulStop()

	// 将server结构体注册到grpc服务中
	pb.RegisterHelloServerServer(srv, &server{})
	addr := fmt.Sprintf("%s:%d", host, *Port)
	fmt.Printf("server addr:%s\n", addr)
	go etcdservice.Register(*EtcdAddr, *ServiceName, addr, 5)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		etcdservice.UnRegister(*ServiceName, addr)

		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}

	}()

	// 监听服务
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("监听异常：", err)
		return
	}
}
