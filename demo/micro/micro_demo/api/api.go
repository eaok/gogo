package main

import (
	"context"
	"encoding/json"
	"log"
	proto "micro_demo/api/proto"
	"strings"

	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	micro "github.com/micro/go-micro/v2"
)

type Foo struct{}
type Example struct{}

//使用官方的api.Request api.Response
func (f *Foo) Bar(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Printf("%+v 收到了一条请求\n", req.Method)
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": "got your request" + strings.Join(name.Values, " "),
	})
	rsp.Body = string(b)

	return nil
}

func (e *Example) Call(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	// parse values from the get request
	name := req.Name

	rsp.Msg = string("OK" + name)
	log.Printf("%+v 收到一条Call的调用信息\n", rsp.Msg)

	return nil
}

func main() {
	// Create a new service
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)

	service.Init()

	// Register handler
	proto.RegisterExampleHandler(service.Server(), new(Example))
	proto.RegisterFooHandler(service.Server(), new(Foo))

	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
