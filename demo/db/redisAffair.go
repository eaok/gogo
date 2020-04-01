package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()
	fmt.Println("connect success ...")

	//MULTI
	conn.Send("MULTI")
	conn.Send("INCR", "foo")
	conn.Send("INCR", "bar")
	reply, err := conn.Do("EXEC")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)

	// WATCH
	conn.Do("SET", "xxoo", 10)
	conn.Do("WATCH", "xxoo")
	v, _ := redis.Int64(conn.Do("GET", "xxoo"))
	v = v + 1 // 这里可以基于值做一些判断逻辑
	conn.Do("SET", "xxoo", 100)
	conn.Send("MULTI")
	conn.Send("SET", "xxoo", v)
	r, err := conn.Do("EXEC")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
