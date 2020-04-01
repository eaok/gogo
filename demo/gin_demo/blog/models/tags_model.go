package models

import (
	"blog/dao"
	"strings"
)

// HandleTagsListData 所有标签存到map中
func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		tagList := strings.Split(tag, "&")
		for _, value := range tagList {
			tagsMap[value]++
		}
	}

	return tagsMap
}

// QueryAllTags 查询标签所有的标签
func QueryAllTags() []string {
	var paramList []string

	sqlStr := "select tags from article"
	err := dao.QueryRows(&paramList, sqlStr)
	if err != nil {
		panic(err)
	}

	return paramList
}
