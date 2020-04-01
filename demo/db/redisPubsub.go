package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func Subs() { // 订阅者
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()

	psc := redis.PubSubConn{conn}
	psc.Subscribe("xxxxx") // 订阅channel1频道
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

func Push(message string) { //发布者
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()

	_, err1 := conn.Do("PUBLISH", "xxxxx", message)
	if err1 != nil {
		fmt.Println("pub err: ", err1)
		return
	}
}

func main() {
	go Subs()
	go Push("this is lianshi")
	time.Sleep(time.Second * 5)
}
