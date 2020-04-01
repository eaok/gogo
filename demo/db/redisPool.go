package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var Pool redis.Pool

func init() { // init 用于初始化一些参数，先于main执行
	Pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379",
				redis.DialConnectTimeout(30*time.Millisecond),
				redis.DialReadTimeout(5*time.Millisecond),
				redis.DialWriteTimeout(5*time.Millisecond))
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			// auth认证
			//if _, err := c.Do("AUTH", "password"); err != nil {
			//	c.Close()
			//	return nil, err
			//}

			// 使用select指令选择数据库
			if _, err := c.Do("SELECT", 0); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},

		// 从连接池取出连接时要做的事
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			// t 当前连接被放回pool的时间
			if time.Since(t) < time.Minute { // 当该连接放回池子不到一分钟
				return nil
			}
			_, err := c.Do("PING") // 放回池子超过一分钟的连接需要执行PING操作
			return err
		},
	}
}

func someFunc() {
	conn := Pool.Get() // 从池子中取出一个连接
	defer conn.Close() // 把连接放回池子

	res, err := conn.Do("HSET", "user2", "name", "lianshi")
	fmt.Println(res, err)
	res1, err := redis.String(conn.Do("HGET", "user2", "name"))
	fmt.Printf("res:%s,error:%v\n", res1, err)
}

func main() {
	someFunc()
}
