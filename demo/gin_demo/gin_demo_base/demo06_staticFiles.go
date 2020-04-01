package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//把images映射到static文件系统
	router.Static("/static", "images")
	tmpl := `<img src="static/strawberry.jpg" alt="" width="400">`
	t := template.New("tmpl.html")
	router.SetHTMLTemplate(template.Must(t.Parse(tmpl)))
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tmpl.html", nil)
	})

	//显示指定⽂件夹下的所有⽂件
	router.StaticFS("/dir", http.Dir("."))
	router.StaticFS("/home", http.Dir("/home/beaver"))

	//显示指定⽂件
	router.StaticFile("/file", "./images/strawberry.jpg")

	router.Run(":8080")
}
