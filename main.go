package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 声明router为指向gin.Engine的指针类型
var router *gin.Engine

// 根据Accept header渲染不同的格式
func render(c *gin.Context, data gin.H, templateName string){
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// 渲染json格式
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// 渲染xml格式
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}



func main()  {
	// 创建默认路由，给router赋值
	router := gin.Default()

	// 加载所有的模板
	router.LoadHTMLGlob("templates/*")

	// 索引处理程序
	router.GET("/", showIndexPage)
	// 单篇文章处理程序
	router.GET("/article/view/:article_id",getArticle)
	// 启动路由
	router.Run()
}
