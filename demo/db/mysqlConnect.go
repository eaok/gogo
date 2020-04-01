package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4")
	if err != nil {
		fmt.Println("创建数据库对象db失败")
		return
	}
	defer db.Close() // db对象创建成功之后才可能调用它的Close方法

	err = db.Ping() // 真正尝试去连数据库
	if err != nil {
		fmt.Printf("连接数据库失败，err:%v\n", err)
		return
	}
	fmt.Println("连接数据库成功。。")
}
