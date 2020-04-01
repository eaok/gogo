package controllers

import (
	"blog/logger"
	"blog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func IndexHandler(c *gin.Context) {
	articleList, err := models.QueryAllArticle()
	if err != nil {
		logger.Error("models.QueryCurrUserArticleWithPage failed", zap.Any("error", err))
	}
	logger.Debug("models.QueryCurrUserArticleWithPage", zap.Any("articleList", articleList))
	c.HTML(http.StatusOK, "index.html", gin.H{"articleList": articleList})
}
