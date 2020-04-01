package dao

import (
	"blog/config"
	"fmt"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitMySQL(cfg *config.MySQLConfig) (err error) {
	logs.Info("InitMySQL....")
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = sqlx.Connect("mysql", str)
	if err != nil {
		return
	}

	if err = CreateTableWithUser(); err != nil {
		return
	}

	if err = CreateTableWithArticle(); err != nil {
		return
	}

	if err = CreateTableWithAlbum(); err != nil {
		return
	}

	return
}

// CreateTableWithUser 创建用户表
func CreateTableWithUser() (err error) {
	sqlStr := `CREATE TABLE IF NOT EXISTS users(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        status INT(4),
        create_time INT(10)
        );`

	_, err = ModifyDB(sqlStr)
	return
}

// CreateTableWithArticle 创建文章表
func CreateTableWithArticle() (err error) {
	sqlStr := `create table if not exists article(
        id int(4) primary key auto_increment not null,
        title varchar(30),
        author varchar(20),
        tags varchar(30),
        short varchar(255),
        content longtext,
        create_time int(10),
        status int(4)
        );`
	_, err = ModifyDB(sqlStr)
	return
}

// CreateTableWithAlbum 创建图片表
func CreateTableWithAlbum() (err error) {
	sqlStr := `create table if not exists album(
        id int(4) primary key auto_increment not null,
        filepath varchar(255),
        filename varchar(64),
        status int(4),
        create_time int(10)
        );`
	_, err = ModifyDB(sqlStr)
	return
}

//操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
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

//查询
func QueryRowDB(dest interface{}, sql string, args ...interface{}) error {
	return db.Get(dest, sql, args...)
}

// 查询多条
func QueryRows(dest interface{}, sql string, args ...interface{}) error {
	return db.Select(dest, sql, args...)
}
