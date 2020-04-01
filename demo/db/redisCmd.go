package main

import (
	"fmt"
	"reflect"

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

	// String get/set
	_, err = conn.Do("SET", "name", "lianshi")
	if err != nil {
		fmt.Println("redis set error:", err)
	}
	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Got name: %s \n", name)
	}

	// String mget/mset
	_, err = conn.Do("MSET", "name", "xiaoxiao", "age", 18)
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.Strings(conn.Do("MGET", "name", "age"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		resType := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", resType)
		fmt.Printf("MGET name: %s \n", res)
		fmt.Println(len(res))
	}

	// 设置过期时间
	_, err = conn.Do("expire", "name", 10) // 10秒过期
	if err != nil {
		fmt.Println("set expire error: ", err)
		return
	}

	// List操作
	_, err = conn.Do("LPUSH", "list1", "ele1", "ele2", "ele3", "ele4")
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	//res, err := redis.String(conn.Do("LPOP", "list1"))//获取栈顶元素
	//res, err := redis.String(conn.Do("LINDEX", "list1", 3)) //获取指定位置的元素
	res, err = redis.Strings(conn.Do("LRANGE", "list1", 0, 3)) //获取指定下标范围的元素
	if err != nil {
		fmt.Println("redis POP error:", err)
	} else {
		resType := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", resType)
		fmt.Printf("res  : %s \n", res)
	}

	// HASH
	//user -> {name: lianshi age:18}
	_, err = conn.Do("HSET", "user", "name", "lianshi", "age", 18)
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res0, err := redis.Int64(conn.Do("HGET", "user", "age"))
	if err != nil {
		fmt.Println("redis HGET error:", err)
	} else {
		resType := reflect.TypeOf(res0)
		fmt.Printf("res type : %s \n", resType)
		fmt.Printf("res  : %d \n", res0)
	}

	// PIPELINE
	conn.Send("HSET", "user1", "name", "lianshi", "age", "30")
	conn.Send("HSET", "user1", "sex", "female")
	conn.Send("HGET", "user1", "age")
	conn.Flush()

	res1, err := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("Receive res2:%v\n", res2)
	res3, err := conn.Receive()
	fmt.Printf("Receive res3:%s\n", res3)
}
