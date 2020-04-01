package controllers

import (
	"blog/logger"
	"blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// HomeGet 获取首页显示的文章
// 可以传入tag/page
func HomeGet(c *gin.Context) {
	//获取session，判断是否登录
	isLogin := c.MustGet("is_login").(bool)
	userName := c.MustGet("login_user").(string)

	tag := c.Query("tag")
	page, _ := strconv.Atoi(c.Query("page"))
	logger.Debug("HomeGet", zap.Int("page", page), zap.String("tag", tag))

	var articleList []*models.Article
	if len(tag) > 0 {
		//按照指定的标签搜索
		articleList, _ = models.QueryArticlesWithTag(tag)
	} else {
		if page <= 0 {
			page = 1
		}
		articleList, _ = models.QueryCurrUserArticleWithPage(userName, page)
	}
	logger.Debug("models.QueryCurrUserArticleWithPage", zap.Any("articleList", articleList))

	//所有文章在后端渲染出来HTML数据
	data := models.GenHomeBlocks(articleList, isLogin)
	pageData := models.GenHomePagination(page)
	logger.Debug("home.html", zap.Any("data", data), zap.Any("data", data), zap.Any("pageData", pageData))
	c.HTML(http.StatusOK, "home.html", gin.H{"isLogin": isLogin, "data": data, "pageData": pageData})
}
