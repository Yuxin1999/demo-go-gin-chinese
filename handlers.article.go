package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context){
	articles := getAllArticles()

	// 根据格式渲染
	render(c, gin.H{
			"title": "Home Page",
			"payload": articles}, "index.html")
}

func getArticle(c *gin.Context){
	// 检查文章id是否有效
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil{
		// 检查是否获取了文章
		if article, err := getArticleByID(articleID); err == nil{
			// 若获取了文章，则渲染
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"title": article.Title,
					"payload": article,
				},
				)
		} else {
			// 若未找到该文章，返回该错误
			c.AbortWithError(http.StatusNotFound, err)
		}
	}else {
		// 若文章ID无效，返回该错误
		c.AbortWithStatus(http.StatusNotFound)
	}
}
