package controllers

import (
	"blog/logger"
	"blog/models"
	"blog/utils"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginGet 登录页
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

// LoginPost 登录页表单处理
func LoginPost(c *gin.Context) {
	userName := c.PostForm("username")
	password := c.PostForm("password")
	logger.Info("LoginPost", zap.Any("username", userName), zap.Any("password", password))

	//在数据库中匹配
	id := models.QueryUserWithParam(userName, utils.Sha256(password))
	if id > 0 {
		//返回cookie
		session := sessions.Default(c)
		session.Set("login_user", userName)
		session.Save()

		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}
}

// LogoutHandler 退出登录
func LogoutHandler(c *gin.Context) {
	//清除该用户登录状态的数据
	session := sessions.Default(c)
	session.Delete("login_user")
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}
