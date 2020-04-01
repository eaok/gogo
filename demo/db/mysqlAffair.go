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

	//开启事务
	tx, _ := db.Begin()
	//提供一组sql操作
	var aff1, aff2 int64 = 0, 0
	result1, _ := tx.Exec("UPDATE account SET money=1000 WHERE id=?", 1)
	result2, _ := tx.Exec("UPDATE account SET money=3000 WHERE id=?", 2)
	if result1 != nil {
		aff1, _ = result1.RowsAffected()
	}
	if result2 != nil {
		aff2, _ = result2.RowsAffected()
	}

	if aff1 == 1 && aff2 == 1 {
		//提交事务
		tx.Commit()
		fmt.Println("操作成功。。")
	} else {
		//回滚
		tx.Rollback()
		fmt.Println("操作失败。。。回滚。。")
	}
}
