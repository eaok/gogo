package models

import (
	"blog/dao"
	"blog/logger"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

const (
	pageSize = 4
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime int64 `db:"create_time"`
	Status     int   // Status=0为正常，1为删除，2为冻结
}

// AddArticle 添加文章
func AddArticle(article Article) (int64, error) {
	sqlStr := "insert into article(title,tags,short,content,author,create_time,status) values(?,?,?,?,?,?,?)"
	return dao.ModifyDB(sqlStr, article.Title, article.Tags, article.Short, article.Content, article.Author,
		article.CreateTime, article.Status)
}

// UpdateArticle 更新文章
func UpdateArticle(article *Article) (int64, error) {
	sqlStr := "update article set title=?,tags=?,short=?,content=? where id=?"
	return dao.ModifyDB(sqlStr, article.Title, article.Tags, article.Short, article.Content, article.Id)
}

// DeleteArticle 删除文章
func DeleteArticle(id string) (int64, error) {
	sqlStr := "delete from article where id=?"
	return dao.ModifyDB(sqlStr, id)
}

// QueryCurrUserArticleWithPage 按页查询当前用户的文章
//总共有10条数据，每页显示4条。  --> 总共需要(10-1)/4+1 页。
//问第2页数据是哪些？           --> 5,6,7,8  (2-1)*4,4
func QueryCurrUserArticleWithPage(username string, pageNum int) (articleList []*Article, err error) {
	sqlStr := "select id,title,tags,short,content,author,create_time from article where author=? limit ?,?"
	pageNum--
	args := []interface{}{username, pageNum * pageSize, pageSize}
	if err := dao.QueryRows(&articleList, sqlStr, args...); err != nil {
		logger.Error("QueryCurrUserArticleWithPage, ", zap.Any("error", err))
		return nil, err
	}
	logger.Debug("QueryCurrUserArticleWithPage,", zap.Any("articleList", articleList))
	return articleList, nil
}

// QueryAllArticle 查询所有文章
func QueryAllArticle() ([]*Article, error) {
	sqlStr := "select id,title,tags,short,content,author,create_time from article"
	var articleList []*Article
	err := dao.QueryRows(&articleList, sqlStr)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}

// QueryArticleWithId 根据id查单篇文章
func QueryArticleWithId(id string) (article *Article, err error) {
	article = new(Article)
	sqlStr := "select id,title,tags,short,content,author,create_time from article where id=?"
	err = dao.QueryRowDB(article, sqlStr, id)
	return
}

// QueryArticlesByIds 根据id查文章 按顺序
func QueryArticlesByIds(ids []int64, idStrs []string) ([]*Article, error) {
	sqlStr := "select id, title from article where id in (?) order by FIND_IN_SET(id, ?)"
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(idStrs, ","))
	if err != nil {
		logger.Error("QueryArticlesByIds", zap.Any("error", err))
		return nil, err
	}
	var dest []*Article
	err = dao.QueryRows(&dest, query, args...)
	return dest, err
}

// QueryArticlesWithTag 按照标签查询文章
/*
有四种情况
	1.左右两边有&符和其他符号
	2.左边有&符号和其他符号，同时右边没有任何符号
	3.右边有&符号和其他符号，同时左边没有任何符号
	4.左右两边都没有符号

通过%去匹配任意多个字符，至少是一个
*/
func QueryArticlesWithTag(tag string) ([]*Article, error) {
	sqlStr := "select id,title,tags,short,content,author,create_time from article"
	sqlStr += " where tags like '%&" + tag + "&%'"
	sqlStr += " or tags like '%&" + tag + "'"
	sqlStr += " or tags like '" + tag + "&%'"
	sqlStr += " or tags like '" + tag + "'"
	fmt.Println(sqlStr)

	var dest []*Article
	err := dao.QueryRows(&dest, sqlStr)

	return dest, err
}

// QueryArticleRowNum 查询文章的总条数
func QueryArticleRowNum() (num int, err error) {
	err = dao.QueryRowDB(&num, "select count(id) from article")
	return
}
