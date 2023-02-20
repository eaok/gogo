package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Student struct {
	StudentID int    `db:"student_id"`
	Firstname string `db:"firstname"`
	Lastname  string `db:"lastname"`
	Age       int    `db:"age"`
}

var Db *sqlx.DB

func initDb() (err error) {
	dns := "root:root@tcp(localhost:3306)/mydb"

	//Open 可能只是验证这些参数，并不会去连接database，要验证这个连接是否成功，使用ping()方法
	database, err := sqlx.Open("mysql", dns)
	if err != nil {
		log.Printf("open mysql failed, err:%v ", err)
		return
	}
	if err = database.Ping(); err != nil {
		log.Printf("connet mysql failed, err:%v ", err)
		return
	}

	Db = database
	log.Printf("connect to mysql succ")
	return
}

//testApi /v3/test-api
func testAPI(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ret": 1000,
	})
}

//arithmeticAction /v3/arithmetic/plus/?a=3&b=4
func arithmeticAction(c *gin.Context) {
	version := c.Param("version")
	action := c.Param("action")
	a, _ := strconv.Atoi(c.Query("a"))
	b, _ := strconv.Atoi(c.Query("b"))
	result := a + b

	c.JSON(http.StatusOK, gin.H{
		"ret":     1000,
		"version": version,
		"action":  action,
		"result":  result,
	})
}

func getStudentList() (list []*Student, err error) {
	sql := "select student_id, firstname, lastname, age from student"
	err = Db.Select(&list, sql)
	if err != nil {
		log.Printf("select from mysql failed, err:%v sql:%v", err, sql)
		return
	}

	return
}

func studentList(c *gin.Context) {
	studentList, _ := getStudentList()
	c.HTML(http.StatusOK, "student.html", gin.H{
		"student_list": studentList,
	})
}

func authToken(c *gin.Context) {
	result := []struct {
		Device string `db:"device"`
		IP     string `db:"ip"`
	}{}

	intAuthToken := c.Query("intAuthToken")
	if intAuthToken != "" {
		intAuthToken = "'" + intAuthToken + "'"
	}

	sql := fmt.Sprintf("SELECT device, ip FROM int_auth_token_cache WHERE int_auth_token = %s", intAuthToken)
	err := Db.Select(&result, sql)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"ret": 1000,
		})
		log.Printf("select from mysql failed, err:%v sql:%v", err, sql)
	} else {
		c.String(http.StatusOK, "Welcome %s user from %s!", result[0].Device, result[0].IP)
	}
}

func main() {
	router := gin.Default()
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})

	//group
	v := router.Group("/v:version")
	{
		v.GET("/test-api", testAPI)
		v.GET("/arithmetic/:action", arithmeticAction)
		v.POST("/arithmetic/:action", arithmeticAction)
	}

	// /tutorial/student/list
	err := initDb()
	if err != nil {
		log.Printf("init Db failed, err:%v", err)
	}
	router.LoadHTMLFiles("templates/student.html")
	router.GET("/tutorial/student/list", studentList)

	// /user/sp100032/wallet/self/detail?intAuthToken=yuqbajnnr
	router.POST("/user/:userId/wallet/self/detail", authToken)

	router.Run(":8080")
}
