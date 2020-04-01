package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//1. 异步
	r.GET("/async", func(c *gin.Context) {
		// goroutine 中只能使⽤只读的上下⽂ c.Copy()
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			// 注意使⽤只读上下⽂
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	//2. 同步
	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		// 注意可以使⽤原始上下⽂
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	r.Run(":8080")
}
