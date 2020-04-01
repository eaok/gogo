package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func anyHandler(c *gin.Context) {
	c.String(http.StatusOK, c.Request.URL.Path)
}

func main() {
	router := gin.Default()

	router.Any("/any", anyHandler)
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "not found")
	})

	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	//url参数
	router.GET("/url", func(c *gin.Context) {
		defaultName := c.DefaultQuery("name", "Guest")
		queryName := c.Query("name")

		str := fmt.Sprintf("defaultName=%s queryName=%s ", defaultName, queryName)
		c.String(http.StatusOK, str)
	})

	//表单参数
	router.POST("/form", func(c *gin.Context) {
		type1 := c.DefaultPostForm("type", "alert")
		username := c.PostForm("username")
		password := c.PostForm("password")
		//hobbys := c.PostFormMap("hobby")
		//hobbys := c.QueryArray("hobby")
		hobbys := c.PostFormArray("hobby")

		str := fmt.Sprintf("type1=%s, username=%s, password=%s, hobbys= %v",
			type1, username, password, hobbys)
		c.String(http.StatusOK, str)
	})

	//上传单个文件
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, err := c.FormFile("file")
		if err != nil {
			panic(err)
		}
		log.Println(file.Filename)

		// 上传文件至指定目录
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	//上传多个文件
	router.POST("/uploads", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)

			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err:%s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files ", len(files)))
	})

	//路由组
	v1 := router.Group("/v1")
	{
		v1.GET("/login", anyHandler)
		v1.GET("/submit", anyHandler)
	}

	//路由组嵌套
	shopGroup := router.Group("/v1")
	{
		shopGroup.GET("/login", anyHandler)
		shopGroup.GET("/submit", anyHandler)

		xx := shopGroup.Group("xx")
		{
			xx.GET("/oo", anyHandler)
			xx.GET("/gg", anyHandler)
		}
	}

	err := router.Run()
	if err != nil {
		panic(err)
	}
}
