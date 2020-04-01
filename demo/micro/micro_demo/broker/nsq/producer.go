package main

import (
	"context"
	proto "micro_demo/broker/nsq/proto"
	"time"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/nsq"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.broker.nsq.clien"),
		micro.Broker(nsq.NewBroker(
			broker.Addrs([]string{"127.0.0.1:4150"}...),
		)),
	)
	srv.Init()
	pub := micro.NewPublisher("go.micro.broker.topic.nsq", srv.Client())
	go func() {
		for i := 0; ; i++ {
			time.Sleep(1 * time.Second)

			_ = pub.Publish(context.TODO(), &proto.Msg{
				Id:      int32(i),
				Current: time.Now().Unix(),
				Data:    "This is an User Data",
			})
		}
	}()

	if err := srv.Run(); err != nil {
		log.Fatalf("error occurs: %v", err)
	}
}
