package main

import (
	"api/db"
	"api/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	// init MySQL
	if err := db.InitMySQL(); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
}

func main() {
	defer db.SqlDB.Close()
	r := gin.Default()

	router := routers.SetupRouter(r)
	router.Run()
}
