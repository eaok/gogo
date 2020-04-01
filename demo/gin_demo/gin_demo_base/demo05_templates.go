package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

//让块动作可用
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func main() {
	router := gin.Default()

	//模板函数
	router.SetFuncMap(template.FuncMap{
		"unsafe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	//加载模板文件
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		router.LoadHTMLGlob("templates/*.html")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	//使用不同目录下名称所有的模板
	router.LoadHTMLGlob("templates/*/*")
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "Users",
		})
	})

	//有模板函数的模板
	router.GET("/unsafe", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.HTML(http.StatusOK, "unsafe.html", gin.H{
			"content": "<a href='https://www.lianshiclass.com'>lianshi</a>",
		})
	})

	//使用html/template渲染模板
	router.GET("/template", func(c *gin.Context) {
		html := template.Must(template.ParseFiles("templates/posts/index.html"))
		router.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "Posts",
		})
	})

	// 模板继承
	router.HTMLRender = loadTemplates("./templates")
	router.GET("/base/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "/base/index",
		})
	})
	router.GET("/base/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"content": "/base/home",
		})
	})

	router.Run()
}
