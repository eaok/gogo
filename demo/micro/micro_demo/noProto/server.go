package main

import (
	"context"
	"log"

	micro "github.com/micro/go-micro/v2"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	print("Hello\n")
	*msg = "你好，" + *name + "回复字串"

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("service.greeter"),
	)
	service.Init()

	// Register handler
	micro.RegisterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
