package main

import "errors"

type article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

// 使用直接赋值初始化文章列表
var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// 创建返回所有文章列表的函数
func getAllArticles() []article{
	return articleList
}

// 通过ID获取文章
func getArticleByID(id int) (*article, error){
	for _, a := range articleList{
		if a.ID == id{
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}