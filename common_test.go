package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// 声明用于储存的临时列表
var tmpArticleList []article

// 在执行测试功能之前，先进行一些设置
func TestMain(m *testing.M){
	// 将gin设置为Test模式
	gin.SetMode(gin.TestMode)
	// 执行其他的测试
	os.Exit(m.Run())
}

// 用于在测试期间创建路由
// 此处创建的路由为该项目标准路由
func getRouter(withTemplates bool) *gin.Engine{
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

// 执行request并测试其reponse
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request,
	f func(w *httptest.ResponseRecorder) bool) {

	// 创建一个记录reponse的recorder
	w := httptest.NewRecorder()

	// 使用输入的路由，执行输入的request，并使用刚刚创建的recorder记录
	r.ServeHTTP(w, req)

	// 将该响应输入布尔函数，判断是否符合布尔函数规则
	if !f(w) {
		t.Fail()
	}
}

// 将一次测试的数据储存在临时列表中
func saveLists() {
	tmpArticleList = articleList
}

// 从临时列表中恢复主列表
func restoreLists() {
	articleList = tmpArticleList
}

