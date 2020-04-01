package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AboutMeGet(c *gin.Context) {
	//获取session
	isLogin := c.GetBool("is_login")

	c.HTML(http.StatusOK, "aboutme.html", gin.H{"isLogin": isLogin, "wechat": "微信：hello", "qq": "QQ：123456", "tel": "Tel：13900000000"})
}
