package routers

import (
	"blog/controllers"
	"blog/logger"
	"blog/middlewares"
	"blog/utils"
	"html/template"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.Engine {
	// 映射模板函数
	router.SetFuncMap(template.FuncMap{
		"timeStr": utils.SwitchTimeStampToStr,
	})

	router.Static("static", "static") //映射静态文件
	router.LoadHTMLGlob("views/*")    //载入模板

	// 载入中间件
	router.Use(middlewares.GinLogger(logger.Logger), middlewares.GinRecovery(logger.Logger, true))
	router.Use(middlewares.Session())

	// 登录注册 无需认证
	{
		router.GET("/register", controllers.RegisterGet)
		router.POST("/register", controllers.RegisterPost)

		router.GET("/login", controllers.LoginGet)
		router.POST("/login", controllers.LoginPost)

		// topN
		router.GET("/article/top/:n", controllers.ArticleTopN)
	}

	// 路由组注册认证中间件
	basicAuthGroup := router.Group("/", middlewares.BasicAuth())
	{
		basicAuthGroup.GET("/home", controllers.HomeGet)
		basicAuthGroup.GET("/", controllers.IndexHandler)
		basicAuthGroup.GET("/logout", controllers.LogoutHandler)

		//路由组
		article := basicAuthGroup.Group("/article")
		{
			article.GET("/add", controllers.AddArticleGet)
			article.POST("/add", controllers.AddArticlePost)

			// 文章详情
			article.GET("/show/:id", controllers.ShowArticleGet)

			// 更新文章
			article.GET("/update", controllers.UpdateArticleGet)
			article.POST("/update", controllers.UpdateArticlePost)

			// 删除文章
			article.GET("/delete", controllers.DeleteArticle)

		}

		// 相册
		basicAuthGroup.GET("/album", controllers.AlbumGet)

		// 文件上传
		basicAuthGroup.POST("/upload", controllers.UploadPost)

		//标签
		basicAuthGroup.GET("/tags", controllers.TagsGet)

		//关于我
		basicAuthGroup.GET("/aboutme", controllers.AboutMeGet)
	}

	return router
}
