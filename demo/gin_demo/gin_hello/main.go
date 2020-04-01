package main

import (
	initRouter "gin_hello/initRouter"
)

func main() {
	router := initRouter.SetupRouter()

	err := router.Run()
	if err != nil {
		panic(err)
	}
}

