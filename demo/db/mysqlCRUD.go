package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func showResult(result sql.Result) {
	fmt.Println(result.LastInsertId()) // 最后插入的数据的id
	fmt.Println(result.RowsAffected()) // 返回受影响的行数
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4")
	if err != nil {
		log.Fatal("connect database failed: ", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("ping failed: ", err)
	}
	fmt.Println("连接数据库成功。。")

	// db 数据库连接池
	db.SetMaxOpenConns(500) // 最大连接数
	db.SetMaxIdleConns(100) // 最大空闲连接数据，业务低峰期保留多少连接
	//db.SetConnMaxLifetime(30 * time.Second)  // 连接空闲时间

	// 1.1 直接插入数据
	result, err := db.Exec("INSERT INTO userinfo(username,departname,created) values(?,?,?)", "杨超越", "技术部", "2019-11-21")
	if err != nil {
		fmt.Printf("插入数据失败:%v\n", err)
		return
	}
	showResult(result)

	// 1.2 通过预编译方式插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) values(?,?,?)")
	if err != nil {
		fmt.Printf("prepare 失败，err:%v\n", err)
		return
	}
	//补充完整sql语句，并执行
	result, err = stmt.Exec("程潇", "体操部", "2020-02-29")
	if err != nil {
		fmt.Printf("插入数据失败,err:%v\n", err)
		return
	}
	showResult(result)
	result, _ = stmt.Exec("五月", "人事部", "2019-11-11") //再次插入数据

	// 3. 更新数据
	result, err = db.Exec("UPDATE userinfo SET username = ?, departname = ? WHERE uid = ?", "法师", "测试部", 7)
	if err != nil {
		fmt.Println("更新数据失败。。", err)
		return
	}
	showResult(result)

	// 4.1 查询
	var username sql.NullString // 数据库的值如果想知道是不是有值 应该使用Null
	var departname, created string
	//row := db.QueryRow("SELECT username,departname,created FROM userinfo WHERE uid=?", 2)
	//err = row.Scan(&username, &departname, &created)
	err = db.QueryRow("SELECT username,departname,created FROM userinfo WHERE uid=?", 8).Scan(&username, &departname, &created)
	if err != nil {
		fmt.Printf("QueryRow failed, err:%v\n", err)
		return
	}
	fmt.Println(username, departname, created)

	// 4.2 查询多条
	sqlStr := `select username, departname, created from userinfo where uid > ?`
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("Query failed, err:%v\n", err)
		return
	}
	defer rows.Close() // 释放rows占用的数据库连接
	for rows.Next() {
		err = rows.Scan(&username, &departname, &created)
		if err != nil {
			fmt.Printf("Scan failed, err:%v\n", err)
			return
		}
		fmt.Println(username, departname, created)
	}
}
