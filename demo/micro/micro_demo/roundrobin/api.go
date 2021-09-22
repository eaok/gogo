package main

import (
	"context"
	"encoding/json"
	"log"
	hello "micro_demo/roundrobin/greeter/proto"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	api "github.com/micro/micro/api/proto"
)

type Say struct {
	Client hello.SayService
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	// 	name, ok := req.Get["name"]
	// 	if !ok || len(name.Values) == 0 {
	// 		return errors.BadRequest("go.micro.api.greeter", "Name cannot be blank hahaha ")
	// 	}

	response, err := s.Client.Hello(ctx, &hello.Request{
		Name: "Allen",
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	//     b, _ := json.Marshal(map[string]string{
	// 		"message": "OK",
	// 	})
	rsp.Body = string(b)

	return nil
}

func main() {
	wrapper := roundrobin.NewClientWrapper()

	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
		micro.WrapClient(wrapper),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Say{Client: hello.NewSayService("go.micro.srv.greeter", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
