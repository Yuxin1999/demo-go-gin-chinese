package main

import "testing"

// 测试函数，测试能否获取所有的文章
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// 检查获取的列表长度是否和原始列表长度相同
	if len(alist) != len(articleList) {
		t.Fail()
	}
	// 检查里面的对象是否都完全相同
	for i, v := range alist{
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title{
			t.Fail()
			break
		}
	}
}
