package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // init()
	"github.com/jmoiron/sqlx"          // database/sql的升级版
)

type User struct {
	Uid        int
	Username   string
	Departname string
	Created    string
}

func main() {
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4")
	if err != nil {
		fmt.Println("创建数据库对象db失败", err)
		return
	}
	defer db.Close() // db对象创建成功之后才可能调用它的Close方法
	err = db.Ping()  // 真正尝试去连数据库
	if err != nil {
		fmt.Printf("连接数据库失败，err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功。。")

	// sqlx查询所有
	var users []User
	err = db.Select(&users, "SELECT uid,username,departname,created FROM userinfo")
	if err != nil {
		fmt.Println("Select error", err)
	}
	fmt.Printf("this is Select res:%v\n", users)

	// sqlx查询单条
	var user User
	err1 := db.Get(&user, "SELECT uid,username,departname,created FROM userinfo where uid = ?", 9)
	if err1 != nil {
		fmt.Println("GET error :", err1)
	} else {
		fmt.Printf("this is GET res:%v", user)
	}

	// sql注入
	name := "yangchaoyue' or 1=1#"
	sqlStr := fmt.Sprintf("select uid,username,departname,created FROM userinfo where username='%s'", name)
	fmt.Println(sqlStr)
	var users1 []User
	err = db.Select(&users1, sqlStr) // 使用自己拼接SQL去查询数据库
	if err != nil {
		fmt.Println("Select error", err)
	}
	fmt.Printf("res:%v\n", users1)
}
