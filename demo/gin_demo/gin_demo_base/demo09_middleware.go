package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//中间件函数，设置参数并统计执行时间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("before middleware...")
		//设置request变量到Context的Key中,通过Get等函数可以取得
		c.Set("request", "client_request")

		//发送request之前
		c.Next()
		//发送request之后

		// 这个c.Write是ResponseWriter,我们可以获得状态等信息
		status := c.Writer.Status()
		fmt.Println("after middleware,", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

//中间件函数，带标志参数
func MiddleWare2(debug bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if debug {

		} else {

		}
		t := time.Now()
		fmt.Println("before middleware")
		c.Set("request", "client_request")

		c.Next()

		status := c.Writer.Status()
		fmt.Println("after middleware,", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func func1(c *gin.Context) {
	c.String(200, "1\n")
	c.Next()
	c.String(200, "11\n")
}

func func2(c *gin.Context) {
	c.String(200, "2\n")
	c.Abort()
	c.String(200, "22\n")
}

func func3(c *gin.Context) {
	c.String(200, "3\n")
	c.Next()
	c.String(200, "33\n")
}

func main() {
	//r := gin.Default()
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	// 可以为每个路由添加任意数量的中间件。
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	//单个路由的中间件
	r.GET("/before", MiddleWare2(false), func(c *gin.Context) {
		request := c.MustGet("request").(string)
		fmt.Println("request:", request)
		c.JSON(http.StatusOK, gin.H{
			"middle_request": request,
		})
	})

	//为组路由添加中间件
	//v1 := r.Group("/v1", MiddleWare())
	v1 := r.Group("/v1")
	v1.Use(MiddleWare())
	{
		v1.GET("/middleware", func(c *gin.Context) {
			//获取gin上下文中的变量
			request := c.MustGet("request").(string)
			req, _ := c.Get("request")
			fmt.Println("request:", request, req)

			c.JSON(http.StatusOK, gin.H{
				"middle_request": request,
				"request":        req,
			})
		})
	}

	//中间件执行流程
	r.GET("flow", func1, func2, func3)

	//全局中间件
	r.Use(MiddleWare())
	r.GET("/router", func(c *gin.Context) {
		request := c.MustGet("request").(string)
		fmt.Println("request:", request)

		c.JSON(http.StatusOK, gin.H{
			"middle_request": request,
		})
	})

	r.Run(":8080")
}
