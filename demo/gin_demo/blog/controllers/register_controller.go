package controllers

import (
	"blog/logger"
	"blog/models"
	"blog/utils"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func RegisterPost(c *gin.Context) {
	//获取表单信息
	userName := c.PostForm("username")
	password := c.PostForm("password")
	rePassword := c.PostForm("repassword")
	logger.Info("RegisterPost", zap.String("userName", userName),
		zap.String("password", password), zap.String("rePassword", rePassword))

	//判断该用户名是否已经被注册
	id := models.QueryUserWithUsername(userName)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户已经存在"})
		return
	}
	//存储注册的用户名和密码
	user := models.User{
		Username:   userName,
		Password:   utils.Sha256(password),
		Status:     0,
		CreateTime: time.Now().Unix(),
	}
	_, err := models.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "注册成功"})
	}
}
