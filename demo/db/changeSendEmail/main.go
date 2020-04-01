package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"

)

type User struct {
	Name string
	Age int
	Country string
	gorm.Model
}

func (*User) BeforeUpdate() (err error) {
	fmt.Println("before")

	e := &email.Email{
		From: "wang <kcoewoys@aliyun.com>",
		To: []string{"kcoewoys@qq.com"},
		Subject: "note: mysql updated",
		Text: []byte("Mysql update"),
	}

	auth := smtp.PlainAuth("", "kcoewoys@aliyun.com", "xxx", "smtp.aliyun.com")
	err = e.Send("smtp.aliyun.com:25", auth)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//如果表不存在就创建表
	if db.HasTable(User{}) == false {
		db.CreateTable(&User{})
	}
	db.AutoMigrate(&User{})

	//插入数据
	user:= &User{
		Name: "xiaozhu",
		Age: 11,
		Country: "China",
	}
	db.Create(user)

	//查询id为1的数据
	db.First(&user, 1)
	fmt.Println(user)


	//更新内容
	var user1 User
	db.Model(&user1).Where("name = ?", "xiaozhu").Update("Age", 12)
	fmt.Println("update User: ", user1)
}