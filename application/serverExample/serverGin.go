package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello index")
}

func main() {
	router := gin.Default()

	router.GET("/", indexHandler)

	router.Run(":8080")
}
