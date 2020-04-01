package main

import (
	"fmt"
	"protobuf/proto/person"

	"github.com/golang/protobuf/proto"
)

func main() {
	p := &person.Person{
		Name: "Jack",
		Age:  10,
		From: "China",
	}
	fmt.Println("原始数据:", p)

	// 序列化
	dataMarshal, err := proto.Marshal(p)
	if err != nil {
		fmt.Println("proto.Unmarshal.Err: ", err)
		return
	}
	fmt.Println("编码数据:", dataMarshal)

	// 反序列化
	entity := person.Person{}
	err = proto.Unmarshal(dataMarshal, &entity)
	if err != nil {
		fmt.Println("proto.Unmarshal.Err: ", err)
		return
	}

	fmt.Printf("解码数据: 姓名：%s 年龄：%d 国籍：%s\n",
		entity.GetName(), entity.GetAge(), entity.GetFrom())
}
