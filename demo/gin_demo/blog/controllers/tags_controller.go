package controllers

import (
	"blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TagsGet(c *gin.Context) {
	isLogin := c.GetBool("is_login")

	tags := models.QueryAllTags()
	tagsMap := models.HandleTagsListData(tags)

	//返回html
	c.HTML(http.StatusOK, "tags.html", gin.H{"isLogin": isLogin, "Tags": tagsMap})
}
