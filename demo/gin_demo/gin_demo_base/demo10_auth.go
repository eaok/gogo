package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 模拟私有数据
var secrets = gin.H{
	"pan": gin.H{"email": "pan@gmail.com", "phone": "123456"},
}

func main() {
	r := gin.Default()

	// 使用 gin.BasicAuth 中间件，设置授权用户
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"pan":    "123456",
		"golang": "123",
	}))

	// 定义路由
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取提交的用户名（AuthUserKey）
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	r.Run(":8080")
}
