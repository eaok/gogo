package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func InitMySQL() (err error) {
	sqlStr := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4"
	//sqlStr := "root:root@tcp(mysql:3306)/test?charset=utf8mb4"
	SqlDB, err = sql.Open("mysql", sqlStr)
	if err != nil {
		return
	}

	if err = SqlDB.Ping(); err != nil {
		return
	}

	if err = CreateTableWithUser(); err != nil {
		return
	}

	return
}

// CreateTableWithUser 创建用户表
func CreateTableWithUser() (err error) {
	sqlStr := `CREATE TABLE IF NOT EXISTS users(
  		id int(11) unsigned NOT NULL AUTO_INCREMENT,
  		name varchar(64) DEFAULT NULL,
 		telephone varchar(12) DEFAULT '',
  		PRIMARY KEY (id)
		) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;`

	_, err = ModifyDB(sqlStr)
	return
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := SqlDB.Exec(sql, args...)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return count, nil
}
