package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.默认启动方式
	router1 := gin.Default()
	go router1.Run()

	//2.http启动方式
	router2 := gin.Default()
	go http.ListenAndServe(":8081", router2)

	//3.带配置的http启动方式
	router3 := gin.Default()
	s := &http.Server{
		Addr:           ":8082",
		Handler:        router3,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go s.ListenAndServe()

	for true {
	}
}
