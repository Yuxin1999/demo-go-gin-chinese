package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


// 测试对home page的get请求
// 对未认证用户返回http状态码200
func TestShowIndexPageUnauthenticated(t *testing.T){
	// 获取一个默认路由
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// 创建一个request获取路由(method为get，路由为/,body部分无）
	req, _ := http.NewRequest("GET", "/", nil)

	// 调用common_test.go中的testHTTPResponse函数
	// 参数为testing.T, 创建的路由， 创建的request与判断状态的bool函数
	// 该布尔函数判断响应是否符合当前一次定义的规则
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// 测试http状态码是否为200
		statusOK := w.Code == http.StatusOK

		// 测试页面标题是否为主页
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}

// 测试对文章页的get请求响应
func TestArticleUnauthenticated(t *testing.T){
	// 获取一个默认路由
	r := getRouter(true)

	r.GET("/article/view/:article_id", getArticle)

	// 创建一个request获取路由(method为get，路由为/article/view/1）
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	// 调用common_test.go中的testHTTPResponse函数
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// 测试http状态码是否为200
		statusOK := w.Code == http.StatusOK

		// 测试页面标题是否为文章标题
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article 1</title>") > 0

		return statusOK && pageOK
	})
}

// 测试accept head为application/json时，应用程序返回JSON格式的文章列表
func TestArticleListJson(t *testing.T){
	// 获取一个默认路由
	r := getRouter(true)

	r.GET("/", showIndexPage)

	// 创建一个request获取路由(method为get，路由为/）
	req, _ := http.NewRequest("GET", "/", nil)
	// request头中添加格式字段
	req.Header.Add("Accept", "application/json")

	// 调用common_test.go中的testHTTPResponse函数
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// 测试http状态码是否为200
		statusOK := w.Code == http.StatusOK

		// 测试响应是否为json格式
		p, err := ioutil.ReadAll(w.Body)
		if err != nil{
			return false
		}
		var articles []article
		// 将响应内容p储存在articles中
		err = json.Unmarshal(p, &articles)
		// json格式&&文章列表长度>2&&状态正确
		return err == nil && len(articles) >=2 && statusOK
	})
}

// 测试accept head为application/xml时，应用程序返回xml格式的文章
func TestArticleXML(t *testing.T){
	// 获取一个默认路由
	r := getRouter(true)

	r.GET("/article/view/:article_id", getArticle)

	// 创建一个request获取路由(method为get，路由为/）
	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	// request头中添加格式字段
	req.Header.Add("Accept", "application/xml")

	// 调用common_test.go中的testHTTPResponse函数
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// 测试http状态码是否为200
		statusOK := w.Code == http.StatusOK

		// 测试响应是否为json格式
		p, err := ioutil.ReadAll(w.Body)
		if err != nil{
			return false
		}
		var a article
		// 将响应内容p储存在article中
		err = xml.Unmarshal(p, &a)
		// xml格式&&文章ID与标题正确&&状态正确
		return err == nil && a.ID == 1 && len(a.Title) >=0 && statusOK
	})
}